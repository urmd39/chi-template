package infrastructure

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	Database *mongo.Database
	Client   *mongo.Client
}

func NewMongoDatastore() *MongoDBStore {
	var mongoDBStore *MongoDBStore
	var replication = ""
	db, session := connectMongoDatabase(DbUsername, CaFile, CertificateFile, PrivateKeyFile,
		DbAuthMechanism, replication)
	if db != nil && session != nil {
		mongoDBStore = new(MongoDBStore)
		mongoDBStore.Database = db
		mongoDBStore.Client = session
		return mongoDBStore
	}
	log.Fatal("Datastore not create")
	return nil

}

func connectMongoDatabase(dbUserName string, caFile string, certificateFile string, privateKeyFile string, dbAuthMechanism string, replication string) (*mongo.Database, *mongo.Client) {
	var connectOne sync.Once
	var db *mongo.Database
	var session *mongo.Client
	var err error
	connectOne.Do(func() {
		opt := configureConnectionInformation(dbUserName, caFile, certificateFile, privateKeyFile, dbAuthMechanism, replication)
		session, err = mongo.NewClient(opt)
		if err != nil {
			ErrLog.Fatal(err)
		}

		err = session.Connect(context.TODO())
		if err != nil {
			ErrLog.Fatal(err)
		}

		err = session.Ping(context.TODO(), nil)
		if err != nil {
			log.Println(err)
		}
		db = session.Database(DbName)

		// TODO: Kiem tra ket noi: Hien thi danh sach database
		// listDatabase, _ := session.ListDatabaseNames(context.TODO(), bson.M{})
		// fmt.Println(listDatabase)
		// fmt.Println(db.Name())
		// listCollection, _ := db.ListCollectionNames(context.TODO(), bson.M{})
		// fmt.Println(listCollection)
	})

	return db, session
}
func configureConnectionInformation(dbUserName string, caFile string, certificateFile string, privateKeyFile string, dbAuthMechanism string, relication string) *options.ClientOptions {
	// cer, err := tls.LoadX509KeyPair(certificateFile, privateKeyFile)
	// if err != nil {
	// 	ErrLog.Fatal(err)
	// }
	// certs, err := ioutil.ReadFile(caFile)
	// if err != nil {
	// 	ErrLog.Fatal(err)
	// }
	// rootCAs, err := x509.SystemCertPool()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if rootCAs == nil {
	// 	rootCAs = x509.NewCertPool()

	// }
	// if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
	// 	log.Fatal("No certs appended, using system certs only")

	// }
	// config := &tls.Config{
	// 	Certificates: []tls.Certificate{cer},
	// 	RootCAs:      rootCAs,
	// }
	// cert := options.Credential{
	// 	AuthMechanism: dbAuthMechanism,
	// 	//AuthSource: "$external",
	// 	Username: DbName,
	// }
	// rp,err:= readpref.New(readpref.PrimaryMode)
	// if err != nil{
	// 	log.Print(err)
	// }
	opt := options.Client().
		ApplyURI("mongodb://" + DbHostPort).
		// SetTLSConfig(config).
		// SetAuth(cert).SetMaxPoolSize(10).
		// SetReplicaSet(relication).
		SetConnectTimeout(10 * time.Second)

	//SetReadPreference(rp)
	return opt

}
