//go:build !mock_test_off

// This will build if build tag mock_test_off is not parsed and will try to mock
// all tests in _test.go
package zb

import (
	"log"
	"os"
	"testing"

	"github.com/aaabigfish/goex/config"
	"github.com/aaabigfish/goex/mock"
	"github.com/aaabigfish/goex/sharedtestvalues"
)

const mockfile = "../../testdata/http_mock/zb/zb.json"

var mockTests = true

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	err := cfg.LoadConfig("../../testdata/configtest.json", true)
	if err != nil {
		log.Fatal("ZB load config error", err)
	}
	var zbConfig *config.Exchange
	zbConfig, err = cfg.GetExchangeConfig("ZB")
	if err != nil {
		log.Fatal("ZB Setup() init error", err)
	}
	zbConfig.API.AuthenticatedSupport = true
	zbConfig.API.AuthenticatedWebsocketSupport = true
	zbConfig.API.Credentials.Key = apiKey
	zbConfig.API.Credentials.Secret = apiSecret
	z.SkipAuthCheck = true
	z.SetDefaults()
	z.Websocket = sharedtestvalues.NewTestWebsocket()
	err = z.Setup(zbConfig)
	if err != nil {
		log.Fatal("ZB setup error", err)
	}

	serverDetails, newClient, err := mock.NewVCRServer(mockfile)
	if err != nil {
		log.Fatalf("Mock server error %s", err)
	}

	err = z.SetHTTPClient(newClient)
	if err != nil {
		log.Fatalf("Mock server error %s", err)
	}
	endpoints := z.API.Endpoints.GetURLMap()
	for k := range endpoints {
		err = z.API.Endpoints.SetRunning(k, serverDetails)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf(sharedtestvalues.MockTesting,
		z.Name)

	os.Exit(m.Run())
}
