package api

import (
	"errors"
	"github.com/Projector-Solutions/Pharaon-api/tool"
	config "github.com/Projector-Solutions/Pharaon-config/migration"
	"github.com/Projector-Solutions/Pharaon-tools/client"
	"log"
	"net/http"
	"sort"
)

type MigrationService struct {
	name           string
	currentVersion int
	migrations     []Migration
	service        *Service
}

func NewMigrationService(name string, migrations []Migration) *MigrationService {
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return &MigrationService{
		name:           name,
		currentVersion: len(migrations),
		migrations:     migrations,
	}
}

func register(name string) (*Service, error) {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/service", config.Service.HttpConfig.ApiUrl),
	})

	err := hc.WriteJSON(ServiceData{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	code, err := hc.Do(tool.LongTimeoutHttpClient())
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, tool.UnexpectedStatusError
	}

	service := &Service{}
	err = hc.ReadJSON(service)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func migrate(name string, migration Migration) error {
	hc := client.NewClient(&http.Request{
		Method: "POST",
		URL:    client.MustFormatURL("%s/service/%s/apply", config.Service.HttpConfig.ApiUrl, name),
	})

	err := hc.WriteJSON(migration)
	if err != nil {
		return err
	}

	code, err := hc.Do(tool.LongTimeoutHttpClient())
	if err != nil {
		return err
	}

	switch code {
	case 200:
		log.Printf("Migration #%d applied\n", migration.Version)
		return nil
	case 400:
		result := &MigrationResult{}
		err := hc.ReadJSON(result)
		if err != nil {
			return err
		}

		return errors.New(*result.Error)
	}

	return tool.UnexpectedStatusError
}

func (m *MigrationService) Actualize() (int, error) {
	service, err := register(m.name)
	if err != nil {
		return 0, err
	}

	count := 0
	m.service = service
	if m.service.Version < m.currentVersion {
		count = m.currentVersion - m.service.Version

		result := MigrationResult{}
		for i := m.service.Version; i < m.currentVersion; i++ {
			migration := m.migrations[i]
			err = migrate(m.name, migration)
			if err != nil {
				return 0, err
			}
		}

		m.service.Version = result.Version
	}

	return count, nil
}
