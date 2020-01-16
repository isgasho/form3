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

	req, _ := http.NewRequest("GET", url.String(), nil)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
