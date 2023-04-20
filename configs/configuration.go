package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

/*
Returns a Configs struct based on the application.yaml file
@author Christopher Fernandes
@copyright
	&copy; Commonwealth of Australia, 2021-2022.<br>
		<font color="0B8900">For Government Official Use Only</font><br>
*/

type JAMS struct {
	Config struct {
		ali   string `mapstructure:"ali"`
		Kafka struct {
			// kafka bootstrap-server
			bootstrapserver string
			// Kafka client id
			clientid string
		}

		Udp struct {
			port int8 `mapstructure:"port"`
		}
	}
}

var (
	filedirectoryPath  string
	fileNameWithoutExt string
	fileExtension      string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Ltime)
	filedirectoryPath = "../resources/"
	fileNameWithoutExt = "application"
	fileExtension = "yaml"
}

func GetConfigCustom(_filedirectoryPath string, _fileNameWithoutExt string, _fileExtension string) (error, map[string]any) {
	if len(_filedirectoryPath) == 0 || len(_fileNameWithoutExt) == 0 || len(_fileExtension) == 0 {
		log.Fatalln("Argument cannot be empty")
	}
	filedirectoryPath = _filedirectoryPath
	fileNameWithoutExt = _fileNameWithoutExt
	fileExtension = _fileExtension
	return GetConfig()
}

/**
  * if run under the other profile. For example application-dev.yaml
**/
func GetConfigProfile(profile string) (error, map[string]any) {
	if len(profile) == 0 {
		log.Fatalln("Profile does not exist")
	}
	fileNameWithoutExt = fileNameWithoutExt + "-" + profile
	return GetConfig()
}

/*
Reads in the config file and returns the config as a struct
*/
func GetConfig() (error, map[string]any) {
	viper.AutomaticEnv()

	viper.SetConfigName(fileNameWithoutExt)
	viper.SetConfigType(fileExtension)
	viper.AddConfigPath(filedirectoryPath)

	// Config file found and successfully parsed
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalln("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config file was found but another error was produced")
		}
	}
	// viper.SetEnvPrefix("JAMS")
	// viper.BindEnv("jams.kafka.client-id")

	jams := viper.Get("jams")
	jamsm, ok := jams.(map[string]any)
	if !ok {
		log.Fatalln("yaml is not properly formated.")
	}
	return nil, jamsm
}
