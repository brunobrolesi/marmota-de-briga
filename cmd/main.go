package main

import (
	"fmt"

	"github.com/brunobrolesi/marmota-de-briga/cmd/db"
)

func main() {
	db.InitScyllaDb()
	fmt.Println("Marmota pronta pra brigar!")
}
