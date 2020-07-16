package use_case

type CovidTracker interface {
	ListCovidCaseInfo(string) string
	ListCovidDeathsInfo(string) string
}

type CovidNews interface {
	ListLatestNews(string) string
}
