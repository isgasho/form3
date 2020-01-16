package apiclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ListParams are optional parameters used to call the List endpoint
type ListParams struct {
	PageNum, PageSize *int
}

// List accepts optional parameters and lists all accounts
func (a *APIClient) List(params ListParams) (accountList AccountListData, err error) {

	rel := &url.URL{Path: "/v1/organisation/accounts"}
	url := a.BaseURL.ResolveReference(rel)
	url.RawQuery = EncodeOptionalQueryParameters(params)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
		return AccountListData{}, err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return AccountListData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return AccountListData{}, err
	}

	err = json.Unmarshal(body, &accountList)
	if err != nil {
		log.Println(err)
		return AccountListData{}, err
	}
	return accountList, nil
}

//EncodeOptionalQueryParameters takes optional parameters and encodes them into a query string
func EncodeOptionalQueryParameters(params ListParams) (queryString string) {

	query := url.Values{}

	if params.PageNum != nil {
		query.Add("page[number]", string(*params.PageNum))
	}
	if params.PageSize != nil {
		query.Add("page[size]", string(*params.PageSize))
	}

	queryString = query.Encode()

	return queryString
}
