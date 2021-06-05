package service

import (
	config "github.com/Projector-Solutions/Pharaon-config"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"pharaon-migration/data"
)

func InitService() error {
	exists, err := initRepo.Check()
	if err != nil {
		return err
	}

	if !exists {
		log.Println("Creating database...")
		return initRepo.Create()
	}

	log.Println("Database exists")
	return nil
}

func Register(serviceData *data.ServiceData) (*data.Service, error) {
	service, err := serviceRepo.GetByName(serviceData.Name)
	if err != nil {
		return nil, err
	}

	if service == nil {
		cfg := config.Service(serviceData.Name)
		if cfg == nil {
			return nil, data.ServiceNotFound
		}

		err = initRepo.InitService(serviceData.Name, cfg.Database.User, cfg.Database.Password)
		if err != nil {
			return nil, err
		}

		id, _ := uuid.NewUUID()
		service = &data.Service{Id: &id, Name: serviceData.Name}

		err = serviceRepo.Insert(service)
		if err != nil {
			return nil, err
		}
	}

	return service, nil
}

func createMigrationRepository(name string) (*data.MigrationRepository, error) {
	serviceConfig := config.Service(name)
	if serviceConfig == nil {
		return nil, data.ServiceNotFound
	}

	databaseConfig := serviceConfig.Database
	databaseConfig.MaxConnections = 1
	return data.NewMigrationRepository(databaseConfig.PsqlConfig())
}

func migrate(name, script string, versionStep int) (int, error, error) {
	service, err := serviceRepo.GetByName(name)
	if err != nil {
		return -1, nil, err
	}

	if service == nil {
		return -1, nil, data.ServiceNotFound
	}

	repo, err := createMigrationRepository(service.Name)
	if err != nil {
		return -1, nil, err
	}

	defer func() {
		repo.Close()
	}()

	var tx *sqlx.Tx
	tx, err = repo.Execute(script)
	if err != nil {
		return service.Version, err, nil
	}

	version := service.Version + versionStep
	err = serviceRepo.UpdateVersion(name, version)
	if err == nil {
		err = tx.Commit()
	} else {
		err = tx.Rollback()
		return -1, nil, err
	}

	return version, nil, err
}

func Apply(name string, migration data.Migration) (int, error, error) {
	return migrate(name, migration.Up, 1)
}

func Rollback(name string, migration data.Migration) (int, error, error) {
	return migrate(name, migration.Down, -1)
}
