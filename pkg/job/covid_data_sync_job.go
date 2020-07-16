package job

import (
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/data_service"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/repository"
	"log"
	"time"
)

func InitCovidDataSyncJob(covidRepo *repository.CovidInfoRepository) {
	go func(covidRepo *repository.CovidInfoRepository) {
		for {
			time.Sleep(time.Second * 60)
			log.Println("starting covid sync job")
			ct := data_service.NewCovidTrackerDataService()
			information, err := ct.GetSummary()
			if err != nil {
				log.Print("error while getting covid summary", err)
				continue
			}

			covidRepo.UpdateSummary(information)
			log.Println("completed covid sync job")
		}

	} (covidRepo)
}
