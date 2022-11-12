package core

import (
	"fmt"
	adapters "homesensor/adapters/internetspeed"
	"homesensor/components/internetspeed/application"
	"homesensor/shared"
)

type Application struct{}

func NewApplication() shared.Application {
	return &Application{}
}

func (*Application) Run() {
	queryBus := NewQueryBus()

	internetSpeedS := application.NewHsInternetSpeedApplicationService(adapters.NewSpeedTestAdapter())
	internetSpeedQh := application.NewHsInternetSpeedQueryHandler(internetSpeedS)

	queryBus.SetHandler(internetSpeedQh.QueryType(), internetSpeedQh)

	internetSpeed, err := queryBus.Send(internetSpeedQh.QueryType(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", internetSpeed)
}
