package use_case

import (
	"fmt"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/constant"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/repository"
	"log"
)

type CovidTrackerUseCase struct {
	covidRepository *repository.CovidInfoRepository
}

func NewCovidTrackerUseCase(covidRepository *repository.CovidInfoRepository) CovidTracker {
	return &CovidTrackerUseCase{
		covidRepository: covidRepository,
	}
}

func (c *CovidTrackerUseCase) ListCovidDeathsInfo(countryCode string) string {
	switch countryCode {
	case "NEW":
		return fmt.Sprintf(constant.DEATHS_NEW_MESSAGE, "world", c.covidRepository.GetGlobalSummary().NewDeaths)
	case "TOTAL":
		return fmt.Sprintf(constant.DEATHS_TOTAL_MESSAGE, "world", c.covidRepository.GetGlobalSummary().TotalDeaths)
	default:
		summary, err := c.covidRepository.GetSummaryByCountryCode(countryCode)
		if err != nil {
			log.Print("ERROR: getting summary by country code for deaths", err)
			return constant.COUNTRY_CODE_INVALID
		}
		return fmt.Sprintf(constant.DEATHS_TOTAL_MESSAGE, summary.Country, summary.TotalDeaths)
	}
}

func (c *CovidTrackerUseCase) ListCovidCaseInfo(countryCode string) string {
	switch countryCode {
	case "NEW":
		return fmt.Sprintf(constant.CASES_NEW_MESSAGE, "world", c.covidRepository.GetGlobalSummary().NewConfirmed)
	case "TOTAL":
		return fmt.Sprintf(constant.CASES_TOTAL_MESSAGE, "world", c.covidRepository.GetGlobalSummary().TotalConfirmed)
	default:
		summary, err := c.covidRepository.GetSummaryByCountryCode(countryCode)
		if err != nil {
			log.Print("ERROR: getting summary by country code for cases ", err)
			return constant.COUNTRY_CODE_INVALID
		}

		return fmt.Sprintf(constant.CASES_TOTAL_MESSAGE, summary.Country, summary.TotalConfirmed)
	}
}