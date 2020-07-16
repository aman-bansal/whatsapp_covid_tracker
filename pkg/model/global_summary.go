package model

type GlobalCovidInformation struct {
	Global *GlobalSummary `json:"Global"`
	Countries []*CountrySummary `json:"Countries"`
}

//{
//    "NewConfirmed": 100282,
//    "TotalConfirmed": 1162857,
//    "NewDeaths": 5658,
//    "TotalDeaths": 63263,
//    "NewRecovered": 15405,
//    "TotalRecovered": 230845
//  }
type GlobalSummary struct {
	NewConfirmed int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths int `json:"NewDeaths"`
	TotalDeaths int `json:"TotalDeaths"`
	NewRecovered int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}

//{
//      "Country": "ALA Aland Islands",
//      "CountryCode": "AX",
//      "Slug": "ala-aland-islands",
//      "NewConfirmed": 0,
//      "TotalConfirmed": 0,
//      "NewDeaths": 0,
//      "TotalDeaths": 0,
//      "NewRecovered": 0,
//      "TotalRecovered": 0,
//      "Date": "2020-04-05T06:37:00Z"
//    },
type CountrySummary struct {
	Country        string `json:"Country"`
	CountryCode    string `json:"CountryCode"`
	Slug           string `json:"Slug"`
	NewConfirmed   int    `json:"NewConfirmed"`
	TotalConfirmed int    `json:"TotalConfirmed"`
	NewDeaths      int    `json:"NewDeaths"`
	TotalDeaths    int    `json:"TotalDeaths"`
	NewRecovered   int    `json:"NewRecovered"`
	TotalRecovered int    `json:"TotalRecovered"`
	Date           string `json:"Date"`
}
