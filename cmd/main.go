package main

import (
	"gin_learning/internal/api"
)

func main() {
	router := api.NewRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
