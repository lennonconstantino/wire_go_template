package main

import (
	"go-wire/inject"
	"log"
)

func main() {
	appCtx, err := inject.InitializeEvent()
	if err != nil {
		log.Fatalf("failed to initialize app context: %v", err)
	}
	appCtx.Repository.GetSql().GetConnection()
	appCtx.Repository.GetCache().GetConnection()
}
