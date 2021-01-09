package main

import (
	config "dev/interview-craft/configs"
	"dev/interview-craft/products"
	"dev/interview-craft/routing"
	"dev/interview-craft/tools"
	"os"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func main() {
	var cfg config.Config
	err := readCfg(&cfg)
	if err != nil {
		log.Error("main-read-config-error")
	}

	e := echo.New()

	//TODO create MongoClient
	mongoClient := tools.NewMongoClient(cfg)

	// Create product manager
	productManager := products.NewProductManager(mongoClient, cfg.Details)

	//TODO create and attach routes
	router := routing.NewRouter(productManager)
	router.AttachRoutes(e)

	//TODO start a sever
	e.Logger.Fatal(e.Start(":1323"))
}

func readCfg(cfg *config.Config) error {
	log.Info("readcfg-read-config")
	defer log.Info("readcfg-read-config-completed")
	file, err := os.Open("configs/config.yml")
	if err != nil {
		log.Error("readcfg-read-config-error")
		return err
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Error("readcfg-read-config-decode-error")
		return err
	}

	return nil

}
