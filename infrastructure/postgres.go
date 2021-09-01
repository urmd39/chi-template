package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBStore struct {
	Database *gorm.DB
}

func NewPostgresDatastore() *PostgresDBStore {
	var postgresDBStore *PostgresDBStore
	db, err := connectPostgreSQLDatabase(DbHost, DbPort, DbName, DbUsername, DbPassword)
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		postgresDBStore = new(PostgresDBStore)
		postgresDBStore.Database = db
		return postgresDBStore
	}
	log.Fatal("Datastore not create")
	return nil
}

func connectPostgreSQLDatabase(dbHost string, dbPort string, dbName string, dbUserName string, dbPassword string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		dbHost, dbPort, dbUserName, dbPassword, dbName)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	return db, err
}
