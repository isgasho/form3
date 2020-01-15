package accountclient

import (
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
)

// Create registers an existing bank account or creates a new one 
func (c *Client) Create(account AccountData) (response string, statusCode int) {

	url := fmt.Sprintf(baseURL)

	json, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	httpClient := &http.Client{}
	
	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), resp.StatusCode
}
