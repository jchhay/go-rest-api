package api

import (
	"fmt"
	"jchhay/go-rest-api-gin/config"
)

func init() {
	// Setup environment variables
	config.LoadEnv()
	config.SetupConfig()
}

func Run() {

	web := NewRouter()
	fmt.Println("Go API REST Running on port " + config.GetConfig().Server.Port)
	fmt.Println("==================>")
	err := web.Run("localhost:" + config.GetConfig().Server.Port)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return

	}
}
