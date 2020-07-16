package model

type NewsTrack struct {
	News []News
}

type News struct {
	Title string `json:"title"`
	WebUrl string `json:"webUrl"`
}