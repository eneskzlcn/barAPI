package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"log"
	"testing"
)

type Settings struct{
	Host string
	ConsumerName string
	ProviderName string
	PactURL string
	PublishVerificationResults	bool
	FailIfNoPactsFound	bool
	DisableToolValidityCheck bool
	BrokerBaseURL string
	BrokerToken string
	ProviderVersion string
	PactFileWriteMode string
}
func (s * Settings) Init(){
	s.Host = "127.0.0.1"
	s.ConsumerName = "foo"
	s.ProviderName = "bar"
	s.PactURL = "https://eneskzlcn.pactflow.io/pacts/provider/bar/consumer/foo/version/1.0.0"
	s.PublishVerificationResults = true
	s.FailIfNoPactsFound = true
	s.DisableToolValidityCheck = true
	s.BrokerBaseURL = "https://eneskzlcn.pactflow.io"
	s.BrokerToken = "L0IzB6WxiCRX7sEdAQoWlQ"
	s.ProviderVersion = "1.0.0"
	s.PactFileWriteMode = "merge"
}
func TestProvider(t *testing.T){
	port,_ := utils.GetFreePort()

	go StartServer(port)

	settings :=	Settings{}
	settings.Init()

	pact := dsl.Pact{
		Consumer:                 settings.ConsumerName,
		Provider:                 settings.ProviderName,
		Host:                     settings.Host,
		DisableToolValidityCheck: settings.DisableToolValidityCheck,
	}
	emptyStateHandler := func() error{
		return nil
	}
	log.Println(pact.Host)
	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:           fmt.Sprintf("http://%s:%d", settings.Host, port),
		PactURLs:                   []string{settings.PactURL},
		BrokerURL:                  settings.BrokerBaseURL,
		BrokerToken:                settings.BrokerToken,
		FailIfNoPactsFound:         settings.FailIfNoPactsFound,
		PublishVerificationResults: settings.PublishVerificationResults,
		ProviderVersion:            settings.ProviderVersion,
		StateHandlers: 			map[string]types.StateHandler{
			"I get pongs array amounf of sent ping count ": emptyStateHandler,
		},
	}
	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(len(verifyResponses), "pact tests run")

}