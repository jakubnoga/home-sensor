package application

import (
	"homesensor/components/internetspeed/domain"
	"homesensor/components/internetspeed/ports"
)

type InternetSpeedApplicationService interface {
	CheckInternetSpeed() (domain.InternetSpeed, error)
}

type HsInternetSpeedApplicationService struct {
	adapter ports.InternetSpeedPort
}

func NewHsInternetSpeedApplicationService(adapter ports.InternetSpeedPort) InternetSpeedApplicationService {
	return &HsInternetSpeedApplicationService{
		adapter,
	}
}

func (s *HsInternetSpeedApplicationService) CheckInternetSpeed() (domain.InternetSpeed, error) {
	return s.adapter.CheckInternetSpeed()
}


