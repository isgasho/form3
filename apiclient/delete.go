package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
)

// Delete deletes an account
func (a *APIClient) Delete(accountID string, version int) (response string, err error) {

	path := fmt.Sprintf("/account/%s", accountID)

	rel := &url.URL{Path: path}
	deleteURL := a.BaseURL.ResolveReference(rel)

	query := url.Values{}
	query.Add("version", string(version))
	deleteURL.RawQuery = query.Encode()

	req, err := http.NewRequest("DELETE", deleteURL.String(), nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil
}
