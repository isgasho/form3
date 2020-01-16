package apiclient

import (
	"fmt"
	"net/http"
	"log"
)

// Delete deletes an account
func (a *APIClient) Delete(accountID string, version int) (err error){

	url := fmt.Sprintf(a.BaseURL.String()+"/v1/organisation/accounts/%s?version=%d", accountID, version)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	return nil
}
