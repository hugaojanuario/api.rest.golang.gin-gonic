package main

import (
	"github.com/hugaojanuario/api-rest-gin/databases"
	"github.com/hugaojanuario/api-rest-gin/routes"
)

func main() {
	databases.ConectaComBancoDeDados()
	routes.HandleRequests()
}
