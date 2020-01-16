package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
		return err.Error(), err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error(), err
	}

	return string(body), nil
}
