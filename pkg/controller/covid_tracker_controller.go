package controller

import (
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/constant"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/data_service"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/repository"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/use_case"
	"strings"
)

type Controller struct {
	newsTracker use_case.CovidNews
	covidInfoTracker use_case.CovidTracker
}

func NewController(infoRepository *repository.CovidInfoRepository, tracker data_service.CovidNewsTracker) Controller {
	return Controller{
		newsTracker:      use_case.NewCovidNewsUseCase(tracker),
		covidInfoTracker: use_case.NewCovidTrackerUseCase(infoRepository),
	}
}

func (c *Controller) HandleWhatsAppQuery(message string) string {
	params := strings.Split(message, " ")
	if len(params) != 2 {
		return constant.CORRECT_CODE_MESSAGE
	}

	switch params[0] {
	case "CASES":
		return c.covidInfoTracker.ListCovidCaseInfo(params[1])
	case "DEATHS":
		return c.covidInfoTracker.ListCovidDeathsInfo(params[1])
	case "NEWS":
		return c.newsTracker.ListLatestNews(params[1])
	default:
		return constant.CORRECT_CODE_MESSAGE
	}
}
