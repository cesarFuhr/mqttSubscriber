package app

import "github.com/cesarFuhr/mqttSubscriber/internal/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	LogStatus command.LogStatusHandler
	StorePIDs command.StorePIDsHandler
	StoreDTCs command.StoreDTCHandler
}
