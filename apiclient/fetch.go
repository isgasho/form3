package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
)

// Fetch gets a single account using the accountID
func (a *APIClient) Fetch(accountID string) (response string, err error) {

	path := fmt.Sprintf("/account/%s", accountID)
	rel := &url.URL{Path: path}
	url := a.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
		return err.Error(), err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return err.Error(), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err.Error(), err
	}

	return string(body), nil
}
