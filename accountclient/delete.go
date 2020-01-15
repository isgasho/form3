package accountclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Delete deletes an account 
func (c *Client) Delete(accountID string, version int) (response string, statusCode int) {

	url := fmt.Sprintf(baseURL+"/%s?version=%d", accountID, version)

	req, _ := http.NewRequest("DELETE", url, nil)

	httpClient := &http.Client{}
	
	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), resp.StatusCode
}
