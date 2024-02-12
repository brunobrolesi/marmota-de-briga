package main

import (
	"github.com/brunobrolesi/marmota-de-briga/cmd/db"
	"github.com/brunobrolesi/marmota-de-briga/internal/infrastructure/api"
)

func main() {
	db.InitScyllaDb()
	api.StartApp()
}
