package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Fetch gets a single account using the accountID
func (a *APIClient) Fetch(accountID string) (response string) {

	path := fmt.Sprintf("/account/%s", accountID)
	rel := &url.URL{Path: path}
	url := a.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(body)
}
