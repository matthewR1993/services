package main

import (
	"os"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/matthewR1993/services/endpoints"
	"github.com/matthewR1993/services/service"
	db "github.com/matthewR1993/services/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v9"
	"github.com/matthewR1993/services/valid"
	
)

var debug bool

func init() {
	// parse console parameters
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()
	
	// setup log settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
}

func main() {
	// setup validator
	valid.Validate = validator.New()

	// Setup log file
	f, errf := os.OpenFile("logfile.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Fatal(errf)
	}
	defer f.Close()
	log.SetOutput(f)

	if debug == true { log.SetOutput(os.Stdout) }

	// Db connection initialization
	var err error
	db.DBCon, err = gorm.Open("postgres", "host=localhost user=matt dbname=storage sslmode=disable password=qwertyuiop")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DBCon.Close()

	// Server initialization
	srv := service.Init()
	srv.SetConf(service.Config{
		Host: "127.0.0.1",
		Port: "8080",
		AssetsDir : "frontend/static",
	})

	// Endpoints registration
	srv.AddEndpoints(
		endpoints.Main,
		endpoints.TestEcho,
		endpoints.RegisterNewUser,
		endpoints.GenerateAuthToken,
	)

	// Shared middleware registration (affect all endpoints)
	srv.AddMiddlewares(
		//endpoints.RequestConsoleLog,
	)

	if debug == true { srv.AddMiddlewares(endpoints.RequestConsoleLog) }

	srv.Run()
}
