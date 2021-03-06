package main

import (
	"beauty/internal/server"
	"beauty/internal/service/postgres"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"beauty/internal/core/client"
	order "beauty/internal/core/order"
	service "beauty/internal/core/service"
	_ "net/http/pprof"
)

type AppConf struct {
	Server server.Conf           `json:"server"`
	DB     postgres.PostgresConf `json:"db"`
}

func main() {
	log.SetFlags(log.Llongfile)
	//runtime.SetBlockProfileRate(1)
	appConf, err := loadConfig("./")
	if err != nil {
		log.Println("Failed to load config")
		return
	}
	//memDB := memorydb.New()
	pgDB, err := postgres.New(appConf.DB)

	clientController := client.New(pgDB)
	orderController := order.New(pgDB)
	serviceController := service.New(pgDB)
	server := server.New(appConf.Server, &clientController, &orderController, &serviceController)
	server.Run()
}

func loadConfig(path string) (AppConf, error) {
	jsonFile, err := os.Open(path + "/conf.json")
	if err != nil {
		log.Println("config file opening error: ", err)
		return AppConf{}, err
	}
	defer jsonFile.Close()
	byteFileContent, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("config file readin error: ", err)
		return AppConf{}, err
	}
	appConf := AppConf{}
	err = json.Unmarshal(byteFileContent, &appConf)
	if err != nil {
		log.Println("config file decoding error: ", err)
		return AppConf{}, err
	}
	return appConf, nil
}
