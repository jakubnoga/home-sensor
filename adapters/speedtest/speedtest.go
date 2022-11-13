package speedtest

import (
	"homesensor/components/internetspeed/domain"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedtestAdapter struct{}

func NewSpeedTestAdapter() *SpeedtestAdapter {
	return &SpeedtestAdapter{}
}

func (s *SpeedtestAdapter) CheckInternetSpeed() (domain.InternetSpeed, error) {
	iS := domain.InternetSpeed{}

	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return iS, err
	}

	servers, err := speedtest.FetchServers(user)
	if err != nil {
		return iS, err
	}

	targets, err := servers.FindServer([]int{})
	if err != nil {
		return iS, err
	}

	for _, server := range targets {
		server.PingTest()
		server.DownloadTest(false)
		server.UploadTest(false)
		
		iS.Download = server.DLSpeed
		iS.Upload = server.ULSpeed
		iS.Ping = server.Latency.Milliseconds()
		iS.Timestamp = time.Now().Format(time.RFC3339)

		break
	}

	return iS, nil
}
