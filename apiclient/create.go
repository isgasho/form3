package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Create registers an existing bank account or creates a new one
func (a *APIClient) Create(account AccountData) (response string) {

	rel := &url.URL{Path: "/account"}
	url := a.BaseURL.ResolveReference(rel)

	json, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(json))
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	return string(body)
}
