package ports

import (
	"homesensor/components/internetspeed/domain"
)

type CheckInternetSpeedPort interface {
	CheckInternetSpeed() (domain.InternetSpeed, error)
}
