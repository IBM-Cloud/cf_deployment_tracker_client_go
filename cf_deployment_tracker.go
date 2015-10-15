package cf_deployment_tracker

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
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

type Event struct {
	DateSent           string   `json:date_sent"`
	CodeVersion        string   `json:code_version"`
	RepositoryURL      string   `json:repository_url"`
<<<<<<< HEAD
	ApplicationName    string   `json:application_name"`
=======
	ApplicationName     string   `json:application_name"`
>>>>>>> a575112d60fdd666a34bfeabb660b5d474d9bfc5
	SpaceID            string   `json:space_id"`
	ApplicationVersion string   `json:application_version"`
	ApplicatonURIs     []string `json:application_uris"`
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

	event := Event{}
	if info.Repository.Url {
		event.RepositoryURL = info.Repository.Url
	}
	if info.Name {
		event.ApplicationName = info.Name
	}
	if info.Version {
		event.ApplicationVersion = info.Version
	}

	request := gorequest.New()
	_, _, errs := request.Post(deploymentTrackerURL).
		Send(event).
		End()

	if errs != nil {
		return errs
	}

}
