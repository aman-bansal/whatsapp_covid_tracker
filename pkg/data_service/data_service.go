package data_service

import "github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"

type CovidTracker interface {
	GetSummary() model.GlobalCovidInformation
}
