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

	// Demonstrate SQL repository usage
	sqlRepo := appCtx.Repository.GetSql()
	sqlRepo.GetConnection() // Simulate connecting to SQL DB
	sqlRepo.GetSql()        // Simulate a SQL operation

	// Demonstrate Cache repository usage
	cacheRepo := appCtx.Repository.GetCache()
	cacheRepo.GetConnection() // Simulate connecting to cache
	cacheRepo.GetCache()      // Simulate a cache operation

	// Direct access via AppContext fields
	appCtx.SqlRepository.GetConnection()
	appCtx.CacheRepository.GetConnection()
}
