package controller

import (
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/data_service"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/repository"
	"log"
	"strconv"
	"strings"
)

type Controller struct {
	covidRepository *repository.CovidInfoRepository
	covidNewsTrackerDataService data_service.CovidNewsTracker
}

func NewController(infoRepository *repository.CovidInfoRepository, tracker data_service.CovidNewsTracker) Controller {
	return Controller{
		covidRepository:             infoRepository,
		covidNewsTrackerDataService: tracker,
	}
}
func (c *Controller) HandleWhatsAppQuery(message string) string {
	switch message {
	case "CASES TOTAL":
		return "Hi there! Total Number of Cases Reported over the world stands at " + strconv.Itoa(c.covidRepository.GetGlobalSummary().TotalConfirmed)
	case "DEATHS TOTAL":
		return "Hi there! Total Number of Deaths Reported over the world stands at " + strconv.Itoa(c.covidRepository.GetGlobalSummary().TotalDeaths)
	case "CASES NEW":
		return "Hi there! New Number of Cases Reported today over the world stands at " + strconv.Itoa(c.covidRepository.GetGlobalSummary().NewConfirmed)
	case "DEATHS NEW":
		return "Hi there! New Number of Deaths Reported today over the world stands at " + strconv.Itoa(c.covidRepository.GetGlobalSummary().NewDeaths)
	default:
		params := strings.Split(message, " ")
		if params[0] == "CASES" {
			summary, err := c.covidRepository.GetSummaryByCountryCode(params[1])
			if err != nil {
				log.Print("ERROR: getting summary by country code for cases ", err)
				return "Hi there! Country code provided by you is invalid. Please check again"
			}

			return "Hi there! Number of Cases Reported in " + summary.Country + " Stands at " + strconv.Itoa(summary.TotalConfirmed)
		}

		if params[0] == "DEATHS" {
			summary, err := c.covidRepository.GetSummaryByCountryCode(params[1])
			if err != nil {
				log.Print("ERROR: getting summary by country code for deaths", err)
				return "Hi there! Country code provided by you is invalid. Please check again"
			}

			return "Hi there! Number of Deaths Reported in " + summary.Country + " Stands at " + strconv.Itoa(summary.TotalDeaths)
		}

		if params[0] == "NEWS" {
			if params[1] == "GLOBAL" {
				news, err := c.covidNewsTrackerDataService.GetLatestNews("global")
				if err != nil {
					log.Print("ERROR: getting latest news about covid. Please check your country code.", err)
					return "Hi there! We are facing Issues while getting latest news. We will be up soon. In the meantime do check your country code if its valid."
				}

				return getResponseMessage(news)
			}
			news, err := c.covidNewsTrackerDataService.GetLatestNews(params[1])
			if err != nil {
				log.Print("ERROR: getting latest news about covid. Please check your country code.", err)
				return "Hi there! We are facing Issues while getting latest news. We will be up soon. In the meantime do check your country code if its valid."
			}

			return getResponseMessage(news)
		}

		return "Code provided by you is wrong! Correct formats are \n 1. CASES TOTAL \n 2. DEATHS TOTAL \n " +
			"3. CASES NEW \n 4. DEATHS NEW \n 5. CASES <COUNTRY CODE IN CAPS> \n 6. DEATHS <COUNTRY CODE IN CAPS> \n" +
			"7. NEWS GLOBAL \n 8. NEWS <COUNTRY CODE IN CAPS>"
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
