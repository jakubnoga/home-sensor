package core

import (
	"homesensor/adapters/speedtest"
	"homesensor/components/internetspeed/application"
	"homesensor/core/annotations"
	"homesensor/core/querybus"
	"homesensor/shared"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Application struct{}

func NewApplication() shared.Application {
	return &Application{}
}

func (*Application) Run() {
	fx.New(
		fx.Provide(
			zap.NewExample,
		),
		fx.Provide(
			annotations.AsQueryBus(querybus.NewQueryBus),
			annotations.AsQueryHandlerRegistration(application.NewInternetSpeedQueryHandler),
			annotations.AsInternetSpeedApplicationService(application.NewInternetSpeedApplicationService),
			annotations.AsInternetSpeedPort(speedtest.NewSpeedTestAdapter),
		),
		fx.Invoke(func(log *zap.Logger, b shared.QueryBus) {
			// log.Info("Sending CheckInternetSpeed query")
			// result, err := b.Send("CheckInternetSpeed", "")
			// log.Sugar().Infof("CheckInternetSpeed query result: %+v, error: %+v", result, err)
		}),
	).Run()
}
