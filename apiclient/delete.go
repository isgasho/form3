package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Delete deletes an account
func (a *APIClient) Delete(accountID string, version int) (response string) {

	path := fmt.Sprintf("/account/%s", accountID)

	rel := &url.URL{Path: path}
	deleteURL := a.BaseURL.ResolveReference(rel)

	query := url.Values{}
	query.Add("version", string(version))
	deleteURL.RawQuery = query.Encode()

	req, _ := http.NewRequest("DELETE", deleteURL.String(), nil)

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
