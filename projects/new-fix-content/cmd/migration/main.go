package main

import "new-fix-content/cmd/api"

func main() {
	server := api.NewApiServer(":8080")

	err := server.Run()

	if err != nil {
		panic(err)
	}

	return
}
