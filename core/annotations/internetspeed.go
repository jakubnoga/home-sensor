package annotations

import (
	"homesensor/components/internetspeed/application"
	"homesensor/components/internetspeed/ports"

	"go.uber.org/fx"
)

func AsInternetSpeedPort(p any) any {
	return fx.Annotate(
		p,
		fx.As(new(ports.CheckInternetSpeedPort)),
	)
}

func AsInternetSpeedApplicationService(s any) any {
	return fx.Annotate(
		s,
		fx.As(new(application.InternetSpeedApplicationService)),
	)
}