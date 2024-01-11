package main

import (
	"jchhay/go-rest-api-gin/api"
	_ "jchhay/go-rest-api-gin/docs"
)

// @title Go Rest API
// @version 1
// @description This is a sample rest api server using golang and gin framework
// @contact.name API Support
// @contact.email XXXXXXXXXXXXXXXX
// @host localhost:3000
// @BasePath /api/v1
func main() {

	api.Run()
}
