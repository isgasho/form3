package apiclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
)

// Create registers an existing bank account or creates a new one
func (a *APIClient) Create(account AccountData) (created AccountData, err error) {

	rel := &url.URL{Path: "/v1/organisation/accounts"}
	url := a.BaseURL.ResolveReference(rel)

	jsonPayload, err := json.Marshal(account)
	if err != nil {
		log.Println(err)
		return AccountData{}, err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println(err)
		return AccountData{}, err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return AccountData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return AccountData{}, err
	}

	err = json.Unmarshal(body, &created)
	if err != nil {
		log.Println(err)
		return AccountData{}, err
	}

	return created, nil
}
