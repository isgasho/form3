package accountclient

import (
	"fmt"
	"net/url"
	"io/ioutil"
	"net/http"
)

// List Params are optional parameters used to query
type ListParams struct {
	PageNum, PageSize, FilterAttr, FilterValue string
  }


// List lists all organisations with the ability to filter
// varadic argument used to pass optional parameters
func (c *Client) List(params ListParams) (response string, statusCode int) {

	query := url.Values{}

	if params.PageNum != ""{
		query.Add("page[number]", params.PageNum)
	}
	if params.PageSize != ""{
		query.Add("page[size]", params.PageSize)
	}
	if params.FilterAttr != "" && params.FilterValue != ""{
		key := fmt.Sprintf("filter[%s]", params.FilterAttr)
		query.Add(key, params.FilterValue)
	}

	queryString := query.Encode()

	ListURL := 	fmt.Sprintf(baseURL+"?%s", queryString)

	req, _ := http.NewRequest("GET", ListURL, nil)

	httpClient := &http.Client{}
	
	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), resp.StatusCode
}
