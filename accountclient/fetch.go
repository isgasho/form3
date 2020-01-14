package accountclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetch gets a single account using the accountID
func (c *Client) Fetch(accountID string) (response string, statusCode int) {

	url := fmt.Sprintf(baseURL+"/%s", accountID)

	req, _ := http.NewRequest("GET", url, nil)

	httpClient := &http.Client{}
	
	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), resp.StatusCode
}
