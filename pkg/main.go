package main

import (
	"fmt"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/controller"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/data_service"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/job"
	"github.com/aman-bansal/whatsapp_covid_tracker/pkg/repository"
	"log"
	"net/http"
)

func main() {
	covidRepo, err := repository.NewCovidDataTrackerRepository()
	if err != nil {
		panic("ERROR: initializing covid repository")
	}

	job.InitCovidDataSyncJob(covidRepo)
	tracker, err := data_service.NewCovidNewsTrackerDataService()
	if err != nil {
		log.Print("ERROR: while getting covid news tracker data service")
		panic("ERROR: while getting covid news tracker data service")
	}

	covidHandler := controller.NewController(covidRepo, tracker)
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			_, _ = fmt.Fprint(w, "Hi there! We are having some issue at our end. Please stay tuned. We will be back soon.")
			return
		}

		response := covidHandler.HandleWhatsAppQuery(r.Form.Get("Body"))
		_, _ = fmt.Fprint(w, response)
		return
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}