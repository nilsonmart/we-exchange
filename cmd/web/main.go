package main

import (
	"github.com/nilsonmart/we-exchange/internal/infra"
)

func main() {

	r := router()

	defer infra.InitSQLite()

	err := r.Run(":8080")
	if err != nil {
		//TODO Log error
		panic(err)
	}

}
