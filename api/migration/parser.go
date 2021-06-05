package api

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func LoadMigrations(dir string) ([]Migration, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	result := make([]Migration, 0)
	for _, f := range files {
		version, err := strconv.Atoi(f.Name())
		if err != nil {
			return nil, err
		}

		upScript, err := readVersionScript(dir, version, "up")
		if err != nil {
			return nil, err
		}

		downScript, err := readVersionScript(dir, version, "down")
		if err != nil {
			return nil, err
		}

		migration := Migration{Version: version, Up: upScript, Down: downScript}
		result = append(result, migration)
	}

	return result, nil
}

func readVersionScript(dir string, version int, t string) (string, error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s/%d/%s.sql", dir, version, t))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
