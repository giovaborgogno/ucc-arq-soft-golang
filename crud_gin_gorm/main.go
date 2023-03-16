package main

import (
	"crud_gin_gorm/database/config"
	"crud_gin_gorm/database/migration"
	"crud_gin_gorm/routes"
	"log"
)

func main() {

	if err := migration.Migrate(config.GetDB()); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	r := routes.SetupRouter()
	r.Run(":8080")

}
