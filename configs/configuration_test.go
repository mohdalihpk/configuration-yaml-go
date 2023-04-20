package configs

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {

	err, config := GetConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config["ali"])
}
