package models

type Result string
type Search func(query string) Result


type Cars struct {
	Id			int `json:"id"`
	Name		string `json:"name"`
	Description	string	`json:"description"`
	Range 		int `json:"range"`
	Battery 	int `json:"battery"`
	AWD 		bool `json:"awd"`
}

type Thumbnail struct {
	Id			int `json:"id"`
	Name        string `json:"name"`
	CarImageUrl []string `json:"car_image_url"`
}

type Video struct {
	Id 			int `json:"id"`	
	Name 		string `json:"name"`
	Source 		string `json:"source"`
	VideoURL 	string `json:"video_url"`
}
