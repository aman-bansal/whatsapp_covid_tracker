package data_service

import (
	"encoding/json"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type CovidNewsTrackerDataService struct {
	client http.Client
}

func NewCovidNewsTrackerDataService() (CovidNewsTracker, error){
	return &CovidNewsTrackerDataService{
		client: http.Client{
			Timeout: time.Second * 30,
		},
	}, nil
}

func (c CovidNewsTrackerDataService) GetLatestNews(countryCode string) (*model.NewsTrack, error) {
	log.Print("INFO: Getting Latest News")
	host := "https://api.smartable.ai/coronavirus/news/" + countryCode
	request, err := http.NewRequest("GET", host, nil)
	if err != nil {
		log.Print("ERROR: not able to make the request data", err)
		return nil, err
	}

	request.Header.Set("Subscription-Key", os.Getenv("SMARTABLE_AI_SUBS_KEY"))
	response, err := c.client.Do(request)
	if err != nil {
		log.Print("ERROR: not able to pull the data", err)
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()
	responseData, err := ioutil.ReadAll(response.Body)
	result := new(model.NewsTrack)
	err = json.Unmarshal(responseData, result)
	if err != nil {
		log.Print("ERROR: not able to unmarshal", err)
		return nil, err
	}

	return result, nil
}
