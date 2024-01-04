package api

import (
	"fmt"
)

func Run() {

	web := Setup()
	fmt.Println("Go API REST Running on port " + "8080")
	fmt.Println("==================>")
	_ = web.Run("localhost:8080")
}
