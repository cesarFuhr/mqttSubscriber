package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/adapters"
	"github.com/cesarFuhr/mqttSubscriber/internal/app"
	"github.com/cesarFuhr/mqttSubscriber/internal/app/command"
	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/broker"
	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/config"
	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/database"
	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/exit"
	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/logger"
	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/subscriber"
	"github.com/cesarFuhr/mqttSubscriber/internal/ports"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func main() {
	run()
}

func run() {
	cfg, err := config.LoadConfigs()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	e := make(chan struct{}, 1)
	exit.ListenToExit(e)

	l := logger.NewLogger()
	c := setupMQTTClient(cfg)
	db := setupSQLDatabase(cfg)

	application, appTeardown := newApplication(cfg, l, c, db)

	mqttPort := ports.NewMQTTPort(application)

	subs := subscriber.NewSubscriber(l, c, mqttPort, cfg.Subscriber.Qos)
	if err := subs.ListenAndHandle(); err != nil {
		log.Fatal("Error listening", err)
	}

	gracefullShutdown(ctx, e, appTeardown)
}

func gracefullShutdown(ctx context.Context, e chan struct{}, teardown func()) {
	<-e
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	teardown()
	log.Println("Goodbye...")
}

func newApplication(cfg config.Config, l logger.Logger, c mqtt.Client, db *sql.DB) (app.Application, func()) {
	pidsR := adapters.NewPIDRepository(db)
	dtcsR := adapters.NewDTCRepository(db)

	return app.Application{
			Commands: app.Commands{
				LogStatus: command.NewLogStatusHandler(l),
				StorePIDs: command.NewStorePIDsHandler(&pidsR),
				StoreDTCs: command.NewStoreDTCHandler(&dtcsR),
			},
		}, func() {
			c.Disconnect(1000)
		}
}

func setupMQTTClient(cfg config.Config) mqtt.Client {
	mqttCfg := broker.BrokerCfg{
		Host:          cfg.Broker.Host,
		Port:          cfg.Broker.Port,
		ClientID:      uuid.NewString(),
		AutoReconnect: cfg.Broker.AutoReconnect,
	}

	cli, err := broker.NewBrokerClient(mqttCfg)
	if err != nil {
		panic(err)
	}

	err = broker.Connect(cli)
	if err != nil {
		panic(err)
	}

	return cli
}

func setupSQLDatabase(cfg config.Config) *sql.DB {
	sqlDB, err := database.NewPGDatabase(database.PGConfigs{
		Host:         cfg.Db.Host,
		Port:         cfg.Db.Port,
		User:         cfg.Db.User,
		Password:     cfg.Db.Password,
		Dbname:       cfg.Db.Dbname,
		Driver:       cfg.Db.Driver,
		MaxOpenConns: cfg.Db.MaxOpenConns,
	})
	if err != nil {
		panic(err)
	}

	if err := database.MigrateUp(sqlDB); err != nil {
		panic(err)
	}

	return sqlDB
}
