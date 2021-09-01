package infrastructure

import "os"

func getStringEnvParameter(envParam string, defaultValue string) string {
	if value, ok := os.LookupEnv(envParam); ok {
		return value
	} else {
		return defaultValue
	}
}

const (
	DEVELOP_ENV         = "DEVELOP_ENV"
	API_HOST_PORT       = "API_HOST_PORT"
	HTTP_SWAGGER        = "HTTP_SWAGGER"
	JWT_PUBLIC_KEY_PATH = "JWT_PUBLIC_KEY_PATH"

	DB_HOST      = "DB_HOST"
	DB_PORT      = "DB_PORT"
	DB_HOST_PORT = "DB_HOST_PORT"
	DBNAME       = "DB_DBNAME"
	DB_USERNAME  = "DB_USER"
	DB_PASSWORD  = "DB_PASSWORD"

	CERTIFICATE_FILE  = "CERTIFICATE_FILE"
	PRIVATE_KEY_FILE  = "PRIVATE_KEY_FILE"
	CA_FILE           = "CA_FILE"
	DB_AUTH_MECHANISM = "DB_AUTH_MECHANISM"
)

var (
	DevelopEnv       string
	HostPort         string
	HttpSwagger      string
	JwtPublicKeyPath string
	DbHost           string
	DbPort           string
	DbHostPort       string
	DbName           string
	DbUsername       string
	DbPassword       string
	CertificateFile  string
	PrivateKeyFile   string
	CaFile           string
	DbAuthMechanism  string
)

func loadEnvParameters() {
	DevelopEnv = getStringEnvParameter(DEVELOP_ENV, "dev")
	HostPort = getStringEnvParameter(API_HOST_PORT, ":8080")
	HttpSwagger = getStringEnvParameter(HTTP_SWAGGER, "http://127.0.0.1:8080/api/v1/anfast-avf/swagger/doc.json")
	JwtPublicKeyPath = getStringEnvParameter(JWT_PUBLIC_KEY_PATH, "./infrastructure/config/public.pem")

	DbHost = getStringEnvParameter(DB_HOST, "127.0.0.1")
	DbPort = getStringEnvParameter(DB_PORT, "5432")
	DbHostPort = getStringEnvParameter(DB_HOST_PORT, "127.0.0.1:27017")
	DbName = getStringEnvParameter(DBNAME, "nitrition")
	DbUsername = getStringEnvParameter(DB_USERNAME, "admin")
	DbPassword = getStringEnvParameter(DB_PASSWORD, "")

	// CertificateFile = getStringEnvParameter(CERTIFICATE_FILE, "./infrastructure/config/client.mongo.crt")
	// PrivateKeyFile = getStringEnvParameter(PRIVATE_KEY_FILE, "./infrastructure/config/client.mongo.key")
	// CaFile = getStringEnvParameter(CA_FILE, "./infrastructure/config/mongoCA.crt")
	// DbAuthMechanism = getStringEnvParameter(DB_AUTH_MECHANISM, "MONGODB-X509")
}
