package main

import (
	"bbs-back/internal/server"
	"bbs-back/internal/service/memorydb"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"bbs-back/internal/core/client"
)

type AppConf struct {
	Server server.Conf `json:"server"`
}

func main() {
	log.SetFlags(log.Llongfile)
	appConf, err := loadConfig("./")
	if err != nil {
		log.Println("Failed to load config")
		return
	}
	memDB := memorydb.New()
	ClientController := client.New(&memDB)

	server := server.New(appConf.Server, &ClientController)
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
