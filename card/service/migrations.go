package service

import (
	api "github.com/Projector-Solutions/Pharaon-api/migration"
	"log"
)

func Migrate() error {
	log.Println("Loading migration scripts...")
	migrations, err := api.LoadMigrations("migrations")
	if err != nil {
		return err
	}

	log.Println("Registering and applying migrations...")
	service := api.NewMigrationService("card", migrations)
	count, err := service.Actualize()
	if err != nil {
		log.Println("Migration failed")
		return err
	}

	if count == 0 {
		log.Println("No migration required")
	} else {
		log.Println("Migration success")
	}

	return nil
}
