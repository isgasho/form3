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

	json, _ := json.Marshal(account)

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(json))

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
