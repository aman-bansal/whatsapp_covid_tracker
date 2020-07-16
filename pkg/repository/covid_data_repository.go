package repository

import (
	"errors"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/data_service"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/model"
	"log"
	"sync"
)

type CovidInfoRepository struct {
	covidTracker         data_service.CovidTracker
	mLock                sync.Mutex
	globalSummary        *model.GlobalSummary
	slugVsSummary        map[string]*model.CountrySummary
	codeVsSummary        map[string]*model.CountrySummary
	countryNameVsSummary map[string]*model.CountrySummary
}

func NewCovidDataTrackerRepository() (*CovidInfoRepository, error) {
	ct := data_service.NewCovidTrackerDataService()
	information, err := ct.GetSummary()
	if err != nil {
		log.Print("ERROR: error while getting summary for covid tracker repository", err)
		return nil, err
	}

	slugVsSummary := make(map[string]*model.CountrySummary)
	codeVsSummary := make(map[string]*model.CountrySummary)
	countryNameVsSummary := make(map[string]*model.CountrySummary)
	for _, info := range information.Countries {
		slugVsSummary[info.Slug] = info
		codeVsSummary[info.CountryCode] = info
		countryNameVsSummary[info.Country] = info
	}

	return &CovidInfoRepository{
		covidTracker:         ct,
		mLock:                sync.Mutex{},
		globalSummary:        information.Global,
		slugVsSummary:        slugVsSummary,
		codeVsSummary:        codeVsSummary,
		countryNameVsSummary: countryNameVsSummary,
	}, nil
}

func (c *CovidInfoRepository) GetGlobalSummary() *model.GlobalSummary {
	c.mLock.Lock()
	defer c.mLock.Unlock()
	return c.globalSummary
}

func (c *CovidInfoRepository) GetSummaryByCountryCode(code string) (*model.CountrySummary, error){
	c.mLock.Lock()
	defer c.mLock.Unlock()
	if _, ok := c.codeVsSummary[code]; !ok {
		return nil, errors.New("country not available")
	}
	return c.codeVsSummary[code], nil
}

func (c *CovidInfoRepository) UpdateSummary(information *model.GlobalCovidInformation) {
	c.mLock.Lock()
	defer c.mLock.Unlock()
	for _, info := range information.Countries {
		c.slugVsSummary[info.Slug] = info
		c.codeVsSummary[info.CountryCode] = info
		c.countryNameVsSummary[info.Country] = info
	}
}