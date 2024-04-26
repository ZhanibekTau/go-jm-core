package config

import (
	"encoding/json"
	"fmt"
	"github.com/ZhanibekTau/go-jm-core/pkg/config/structures"
	"github.com/ZhanibekTau/go-jm-core/pkg/helpers/structHelper"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func InitBaseConfig() (*structures.AppConfig, *structures.DbConfig, error) {
	appConfigInterface, err := InitConfig(&structures.AppConfig{HandlerTimeout: 30})
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	appConfig, ok := appConfigInterface.(*structures.AppConfig)
	if !ok {
		log.Fatalf("cannot init app config. Err: %s", ok)
	}

	dbConfigInterface, err := InitConfig(&structures.DbConfig{})
	dbConfig, ok := dbConfigInterface.(*structures.DbConfig)
	if !ok {
		log.Fatalf("cannot init db config. Err: %s", ok)
	}

	return appConfig, dbConfig, nil
}

func InitConfig(config interface{}) (interface{}, error) {
	workingDir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	if _, err := os.Stat(workingDir + "/.env"); err == nil {
		viper.SetConfigFile(workingDir + "/.env")
		viper.SetConfigType("env")
		viper.AutomaticEnv()
		//Find and read the config file
		err := viper.ReadInConfig()

		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}

		viper.Unmarshal(&config)
	}

	envKeys := structHelper.GetFieldsAsUpperSnake(config)

	osEnvMap := make(map[string]string, len(envKeys))

	for _, key := range envKeys {
		if value, exists := os.LookupEnv(key); exists {
			key = strings.ToLower(key)
			osEnvMap[key] = fmt.Sprint(value)
		}
	}

	//	// Convert the map to JSON
	jsonData, _ := json.Marshal(osEnvMap)
	// Convert the JSON to a struct
	json.Unmarshal(jsonData, &config)

	return config, nil
}
