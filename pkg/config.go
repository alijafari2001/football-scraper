package pkg

import (
	"encoding/json"
	"os"
)

func LoadConfigs() (map[string]string, error) {
	var config map[string]string
	paths := []string{"premier-league.json", "la-liga.json", "bundes-liga.json", "serie-a.json", "ligue-1.json"}
	for _, value := range paths {
		filePath := "config/" + value
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &config)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
