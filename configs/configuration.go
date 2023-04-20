package configs

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
)

/*
Returns a Configs struct based on the application.yaml file
@author Christopher Fernandes
@copyright
	&copy; Commonwealth of Australia, 2021-2022.<br>
		<font color="0B8900">For Government Official Use Only</font><br>
*/

type Configs struct {
	Config Configurations
}

type Configurations struct {

	// Unique identifier for truck instance
	Id string
	// Latitude coodinate of truck in decimal degrees
	Lat float64
	// Longitude coodinate of truck in decimal degrees
	Lon float64
	// Heading of truck in degrees
	Hdg float64
	// IP address of the NTP server
	Ntp_address       string
	Ntp_username      string
	Ntp_password      string
	Request_every_sec int
	Kafka_address     string
	Kafka_topic       string
}

/*
Reads in the config file and returns the config as a struct
*/
func GetConfig() Configs {

	// Read in the config file
	confContent, err := ioutil.ReadFile("./resources/application.yaml")
	if err != nil {
		panic(err)
	}

	// Expand environment variables
	confContent = []byte(os.ExpandEnv(string(confContent)))

	viper.SetConfigType("yml")

	var config Configs

	if err := viper.ReadConfig(bytes.NewBuffer(confContent)); err != nil {
		log.Fatalln("Error reading config file", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Unable to decode into struct", err)
	}

	log.Println("Ntp_addrress: ", config.JAMS.GPS.Ntp_address)
	log.Println("Kafka_address: ", config.JAMS.GPS.Kafka_address)
	log.Println("Id: ", config.JAMS.GPS.Id)
	return config
}
