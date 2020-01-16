# Rosie Hamilton

I have written the client library in a separate Go package called `apiclient`

The tests for `apiclient` are in a package named `apiclient_test` this is so that they can interact 
with exposed interface of `apiclient` in the same way that code importing the `apiclient` would 
interact with it, i.e. only access code which has been declared public (starting with capital letters)

I have placed the code for each `apiclient` operation `create`, `fetch`, `list` and `delete` in a separate Go files. 
This is to increase readability and aid debugging i.e. if `Fetch()` caused an error, it would be make sense for the reader to look in `fetch.go`

The core Go team did not set any timeouts on the standard `net/http` client so I have configured the http client to use a sensible timeout of 10 seconds in `client.go`

To help ensure a valid request body is passed to the `create` endpoint I have included a struct representing Account data in `models.go`. I use this struct to marshal values to JSON to form the body of the POST request to the endpoint. I have chosen to expose this struct within the package so that other packages can reuse it for marshalling/unmarshalling.

For the optional parameters which can be passed to the `list` endpoint I chose to store these inside a struct using *int. I chose to use pointers to int as this type can differentiate between 0 and nil

# Testing

BDD was discussed during my phone interview so I have annotated my tests with numbered comments showing the flow of the tests and which parts of my test code satisfy `given`, `when` and `then`.

From the description of the task, I understand the job of `apiclient` is to issue certain calls to the Accounts API and return the responses. The tests for `apiclient` should therefore verify the calls it makes are correct. I have chosen to use a test server from  `net/http/httptest` to intercept the requests made by `apiclient` so that I can verify the calls it makes are correct. This ensures that `apiclient` tests can run independely of the Accounts API that they interact with, which will preventing them from failing if the Accounts API is unavailable. 

For testing the Accounts API itself (the code pre-written by Form3) unit tests asserting on responses generated should exist in the AccountsAPI code base itself and be run by the AccountsAPI pipeline every time the AccountsAPI is built. If this task was to include testing the pre-built Accounts API I would use the acronym BINNEDMATCH (Boundary, Input, Null, Negative, Empty, Documentation, Method, Authorization, Truncate, Content-type, Headers) to generate my test ideas and apply them to the API.


