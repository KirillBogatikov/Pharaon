package psql

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	URL string
	MaxConnections int
}

const (
	urlKey = "DATABASE_URL"
	maxConnectionsKey = "DATABASE_MAX_CONNECTIONS"
)

func ReadFromEnv() (*DatabaseConfig, error) {
	max, err := strconv.Atoi(os.Getenv(maxConnectionsKey))
	if err != nil {
		return nil, err
	}
	if max < 1 {
		return nil, fmt.Errorf("Incorrect %s: %d", maxConnectionsKey, max)
	}
	if max > 20 {
		log.Printf("WARNING: Using max %d connections to postgres!", max)
	}

	url := os.Getenv(urlKey)
	if len(url) == 0 {
		return nil, fmt.Errorf("Empty %s", urlKey)
	}

	return &DatabaseConfig{
		url,
		max,
	}, nil
}