package main

import (
	"nutrition/infrastructure"
	"nutrition/server"
)

// @title Swagger Nitrition API
// @version 1.0
// @description This is list api for nutrition project
// @host 127.0.0.1:8080
// @BasePath /api/v1/nutrition/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

//go:generate ./command.sh ./template/controller.txt ./controller/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/payload.txt ./controller/payload/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/response.txt ./controller/response/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/query.txt ./controller/query/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/service.txt ./service/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/database.txt ./database/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/repository.txt ./database/mongo/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/api.txt ./server/api/$FILENAME.go $CLASS
//go:generate ./command.sh ./template/model.txt ./model/$FILENAME.go $CLASS

func main() {
	server.HandleHttpServer(infrastructure.HostPort)
}
