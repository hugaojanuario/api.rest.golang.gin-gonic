package main

import (
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/http/routes"
)

func main() {
	databases.ConectingInDatabase()
	routes.HandleRequests()
}
