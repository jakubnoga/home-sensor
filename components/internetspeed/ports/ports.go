package ports

import (
	"homesensor/components/internetspeed/domain"
)

type InternetSpeedPort interface {
	CheckInternetSpeed() (domain.InternetSpeed, error)
}