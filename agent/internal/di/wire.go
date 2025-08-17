//go:generate go tool wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Agent Service DI-package
*/
package agent_di

import (
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
)

type AgentService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Metrics       *metrics.Monitoring
	PprofEndpoint profiling.PprofEndpoint
}

// AgentService =========================================================================================================
var AgentSet = wire.NewSet(
	di.DefaultSet,

	NewAgentService,
)

func NewAgentService(
	// Common
	log logger.Logger,
	config *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Observability
	metrics *metrics.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
) (*AgentService, error) {
	return &AgentService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Metrics:       metrics,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,
	}, nil
}

func InitializeAgentService() (*AgentService, func(), error) {
	panic(wire.Build(AgentSet))
}
