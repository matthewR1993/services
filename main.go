package main

import (
	"os"
	"fmt"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/matthewR1993/services/endpoints"
	"github.com/matthewR1993/services/service"
	db "github.com/matthewR1993/services/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v9"
	"github.com/matthewR1993/services/valid"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/matthewR1993/services/redis"
	"github.com/matthewR1993/services/util"	
)

var (
	debug = flag.Bool("debug", false, "Debug mode")
	redisAddress = flag.String("redis-address", ":6379", "Address to the Redis server")
	redisMaxConnections = flag.Int("redis-max-connections", 10, "Max connections to Redis")
	postgresHost = flag.String("pg-host", "127.0.0.1", "postgres ip")
	postgresPort = flag.String("pg-port", "5432", "postgres port")
	appPort = flag.String("app-port", "8080", "app port")
)

func init() {
	// Setup log settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
	
	// parse command line arguments
	flag.Parse()

	// parse config, available from util.AppConf
	util.InitConfig("config.json")
}

func main() {
	// Redis connection pool initialization
	redis.RedisPool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", *redisAddress)
		if err != nil {
			return nil, err
		}
		return c, err
	}, *redisMaxConnections)
	defer redis.RedisPool.Close()

	// Setup validator
	valid.Validate = validator.New()

	// Setup log file
	f, errf := os.OpenFile("logfile.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Fatal(errf)
	}
	defer f.Close()
	log.SetOutput(f)

	if *debug == true { log.SetOutput(os.Stdout) }

	// Db connection initialization
	var err error
	db.DBCon, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=storage sslmode=disable password=%s",
		*postgresHost, *postgresPort, util.AppConf.PgUsername, util.AppConf.PgPassword))
	if err != nil {
		log.Fatal(err)
	}
	defer db.DBCon.Close()

	// Server initialization
	srv := service.Init()
	srv.SetConf(service.Config{
		Host: "127.0.0.1",
		Port: *appPort,
		AssetsDir : "frontend/dist",
	})

	// Endpoints registration
	srv.AddEndpoints(
		endpoints.Main,
		endpoints.TestEcho,
		endpoints.RegisterNewUser,
		endpoints.GenerateAuthToken,
		endpoints.GetUserInformation,
	)

	// Shared middleware registration. It affects all endpoints and it is triggered first.
	srv.AddMiddlewares(
		endpoints.CORSMiddle,
	)

	if *debug == true { srv.AddMiddlewares(endpoints.RequestConsoleLog) }

	srv.Run()
}

