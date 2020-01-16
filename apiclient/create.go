package apiclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Create registers an existing bank account or creates a new one
func (a *APIClient) Create(account AccountData) (response string, err error) {

	rel := &url.URL{Path: "/account"}
	url := a.BaseURL.ResolveReference(rel)

	json, err := json.Marshal(account)
	if err != nil {
		return err.Error(), err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(json))
	if err != nil {
		return err.Error(), err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err.Error(), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error(), err
	}

	return string(body), nil
}
