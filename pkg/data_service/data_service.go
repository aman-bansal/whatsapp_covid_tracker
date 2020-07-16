package data_service

import "github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"

type CovidTracker interface {
	GetSummary() (*model.GlobalCovidInformation, error)
}

type CovidNewsTracker interface {
	GetLatestNews(countryCode string) (*model.NewsTrack, error)
}
