/*
Shortlink application

Auth-service
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
	"github.com/shortlink-org/shortlink/pkg/logger/field"

	agent_di "github.com/shortlink-org/ds/agent/internal/di"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-agent")

	// Init a new service
	service, cleanup, err := agent_di.InitializeAgentService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			service.Log.Error(r.(string)) //nolint:forcetypeassert // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	signal := graceful_shutdown.GracefulShutdown()

	cleanup()

	service.Log.Info("Service stopped", field.Fields{
		"signal": signal.String(),
	})

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // exit code 143 is used to indicate graceful termination
}
