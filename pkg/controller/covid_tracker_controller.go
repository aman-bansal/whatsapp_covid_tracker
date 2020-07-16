package controller

import (
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/repository"
	"log"
	"strconv"
	"strings"
)

type Controller struct {
	covidRepository *repository.CovidInfoRepository
}

func NewController(infoRepository *repository.CovidInfoRepository) Controller {
	return Controller{covidRepository: infoRepository}
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

			return "Hi there! Number of Deaths Reported in " + summary.Country + " Stands at " + strconv.Itoa(summary.TotalConfirmed)
		}

		return "Code provided by you is wrong! Correct formats are \n 1. CASES TOTAL \n 2. DEATHS TOTAL \n " +
			"3. CASES NEW \n 4. DEATHS NEW \n 5. CASES <CODE IN CAPS> \n 6. DEATHS <CODE IN CAPS>"
	}
}