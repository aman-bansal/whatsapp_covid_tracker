package data_service

import (
	"log"
	"testing"
)

func TestCovidTrackerDataService_GetSummary(t *testing.T) {
	service := NewCovidTrackerDataService()
	information, err := service.GetSummary()
	if err != nil {
		t.Fatal("Error while getting summary")
	}

	log.Print(information)
}

func TestCovidNewsTrackerDataService_GetLatestNews(t *testing.T) {
	tracker, err := NewCovidNewsTrackerDataService()
	if err != nil {
		t.Fatal("Error while init news tracker")
	}

	news, err := tracker.GetLatestNews("IN")
	if err != nil {
		t.Fatal("Error while getting latest news")
	}

	log.Print(news)
}