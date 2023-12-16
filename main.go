package main

import (
	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"sample/handler"
	"sample/migrations"
	"sample/store"
)

func main() {
	// Creating GoFr app
	/*gofr.New():function is a constructor provided by the gofr framework.
	It allocates memory for a new gofr application instance and initializes its internal state. */
	//app  becomes the main access point for interaction 
	app := gofr.New()

	// Running migrations - UP
	if err := migration.Migrate("remote-config-data", dbmigration.NewGorm(app.GORM()),
		migrations.All(), dbmigration.UP, app.Logger); err != nil {
		app.Logger.Fatalf("Error in running migrations: %v", err)
	}

	productStore := store.New()
	productHandler := handler.New(productStore)

	// Creating routes
	app.POST("/product", productHandler.Create)
	app.GET("/product/{id}", productHandler.GetByID)
	app.PUT("/product/{id}", productHandler.Update)
	app.DELETE("/product/{id}", productHandler.Delete)

	// Starting server
	app.Start()
}
