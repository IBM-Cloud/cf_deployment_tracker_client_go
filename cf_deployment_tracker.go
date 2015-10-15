package cf_deployment_tracker

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"os"
	"time"
)

var deploymentTrackerURL = "https://deployment-tracker.mybluemix.net/api/v1/track"

type Repository struct {
	Url string
}

type Package struct {
	Name       string
	Version    string
	Repository Repository
}

type App struct {
	DateSent           string   `json:"date_sent"`
	CodeVersion        string   `json:"code_version"`
	RepositoryURL      string   `json:"repository_url"`
	Application_Name   string   `json:"application_name"`
	SpaceID            string   `json:"space_id"`
	ApplicationVersion string   `json:"application_version"`
	ApplicatonURIs     []string `json:"application_uris"`
}

func Track() (errs []error) {
	content, err := ioutil.ReadFile("package.json")
	//exit early if we cant read the file
	if err != nil {
		return
	}

	var info Package
	err = json.Unmarshal(content, &info)
	//exit early if we can't parse the file
	if err != nil {
		return
	}

	var app App
	vcap := os.Getenv("VCAP_APPLICATION")
	err2 := json.Unmarshal([]byte(vcap), &app)
	//exit early if we couldn't read the envar's for the app
	if err2 != nil {
		panic(err2)
	}

	if info.Repository.Url != "" {
		app.RepositoryURL = info.Repository.Url
	}
	if info.Version != "" {
		app.CodeVersion = info.Version
	}

	dateSent := time.Now()
	layout := "2006-01-02T15:04:05.000Z"
	app.DateSent = dateSent.Format(layout)

	request := gorequest.New()
	_, _, errs = request.Post(deploymentTrackerURL).
		Send(app).
		End()

	if errs != nil {
		return errs
	}
	return nil

}
