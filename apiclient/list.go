package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ListParams are optional parameters used to call the List endpoint
type ListParams struct {
	PageNum, PageSize *int
}

// List accepts optional parameters and lists all accounts
func (a *APIClient) List(params ListParams) (response string) {

	rel := &url.URL{Path: "/account"}
	url := a.BaseURL.ResolveReference(rel)
	url.RawQuery = EncodeOptionalQueryParameters(params)

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
