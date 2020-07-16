package use_case

import (
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/constant"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/data_service"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"
	"log"
	"strconv"
)

type CovidNewsUseCase struct {
	covidNewsTrackerDataService data_service.CovidNewsTracker
}

func NewCovidNewsUseCase(tracker data_service.CovidNewsTracker) CovidNews {
	return &CovidNewsUseCase{
		covidNewsTrackerDataService: tracker,
	}
}

func (c *CovidNewsUseCase) ListLatestNews(countryCode string) string {
	switch countryCode {
	case "GLOBAL":
		news, err := c.covidNewsTrackerDataService.GetLatestNews("global")
		if err != nil {
			log.Print("ERROR: getting latest news about covid. Please check your country code.", err)
			return constant.NEWS_ERROR_MESSAGE
		}

		return getResponseMessage(news)
	default:
		news, err := c.covidNewsTrackerDataService.GetLatestNews(countryCode)
		if err != nil {
			log.Print("ERROR: getting latest news about covid. Please check your country code.", err)
			return constant.NEWS_ERROR_MESSAGE
		}

		return getResponseMessage(news)
	}
}

func getResponseMessage(allNews *model.NewsTrack) string {
	message := ""
	count := 1
	for _, news := range allNews.News {
		if count > 10 { break }
		message = message + strconv.Itoa(count) + ". " + news.Title + "\n"
		count = count + 1
	}
	return message
}