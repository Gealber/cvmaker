package maker

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Basic struct {
	Name     string           `json:"name"`
	Label    string           `json:"label"`
	Image    string           `json:"image"`
	Email    string           `json:"email"`
	Phone    string           `json:"phone"`
	URL      string           `json:"url"`
	Summary  []string         `json:"summary"`
	Profiles []NetworkProfile `json:"profiles"`
}

type NetworkProfile struct {
	Network string `json:"network"`
	URL     string `json:"url"`
}

type Education struct {
	Institution string `json:"institution"`
	URL         string `json:"url"`
	Area        string `json:"area"`
	StudyType   string `json:"studyType"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type Work struct {
	Name        string   `json:"name"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Position    string   `json:"position"`
	URL         string   `json:"url"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Summary     string   `json:"summary"`
	HighLights  []string `json:"highlights"`
	Stack       []string `json:"stack"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	HighLights  []string `json:"highlights"`
	Keywords    []string `json:"keywords"`
	Type        string   `json:"type"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type CV struct {
	Basics    Basic       `json:"basics"`
	Education []Education `json:"education"`
	Work      []Work      `json:"work"`
	Projects  []Project   `json:"projects"`
	Languages []Language  `json:"language"`
}

func newCV() *CV {
	body, err := ioutil.ReadFile("resume.json")
	if err != nil {
		log.Fatal("readfile", err)
	}

	var cv CV

	err = json.Unmarshal(body, &cv)
	if err != nil {
		log.Fatal("unmarshal", err)
	}

	return &cv
}
