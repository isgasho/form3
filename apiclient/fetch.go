package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
	"encoding/json"
)

// Fetch gets a single account using the accountID
func (a *APIClient) Fetch(accountID string) (account AccountData, err error) {

	path := fmt.Sprintf("/v1/organisation/accounts/%s", accountID)
	rel := &url.URL{Path: path}
	url := a.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", url.String(), nil)
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

	err = json.Unmarshal(body, &account)
	if err != nil {
		log.Println(err)
		return AccountData{}, err
	}

	return account, nil
}
