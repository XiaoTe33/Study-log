package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	//viper.SetConfigType("yaml")
	//viper.SetConfigFile("config/config.yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	get := viper.Get("mysql")
	fmt.Println(get)

	//viper.SetConfigType("json")
	//viper.SetConfigFile("config/tsconfig.json")
	viper.SetConfigName("tsconfig")
	viper.AddConfigPath("config")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	get = viper.Get("compilerOptions")
	fmt.Println(get)
	get = viper.Get("exclude")
	fmt.Println(get)
	fmt.Println(viper.AllKeys())
}
