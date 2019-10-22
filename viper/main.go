package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	company := viper.GetString("company")
	team := viper.GetString("team")
	auth := viper.GetStringMapString("auth")

	fmt.Println("Empresa: " + company)
	fmt.Println("Team: " + team)
	fmt.Printf("Reading config for auth = %#v\n", auth)

}
