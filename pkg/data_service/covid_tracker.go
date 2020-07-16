package data_service

import (
	"encoding/json"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CovidTrackerDataService struct {
	client http.Client
}

func (c *CovidTrackerDataService) GetSummary() (*model.GlobalCovidInformation, error) {
	log.Print("INFO: Getting Summary")
	host := "https://api.covid19api.com/summary"
	request, err := http.NewRequest("GET", host, nil)
	if err != nil {
		log.Print("ERROR: not able to make the request data", err)
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		log.Print("ERROR: not able to pull the data", err)
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()
	responseData, err := ioutil.ReadAll(response.Body)
	result := new(model.GlobalCovidInformation)
	err = json.Unmarshal(responseData, result)
	if err != nil {
		log.Print("ERROR: not able to unmarshal", err)
		return nil, err
	}

	return result, nil
}

func NewCovidTrackerDataService() CovidTracker {
	return &CovidTrackerDataService{
		client: http.Client{
			Timeout: time.Second * 30,
		},
	}
}
