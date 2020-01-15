# Rosie Hamilton

I have written the client library in a separate Go package called `accountclient`
This is so that the `accountclient` package can be easily imported and reused by other Go program

The tests for `accountclient` are in package named `accountclient_test` this is so that they can only
interact with exposed interface of `accountclient` i.e. only use public methods (starting with capital letters)

I have placed the code for each client operation `create`, `fetch`, `list` and `delete` in a separate file 
This is to increase readability and aid debugging i.e. if `fetch()` caused an error, the dev would know to look in `fetch.go`

The job of `accountclient` is to issue certain calls to the Accounts API, no more and no less. The tests for `accountclient` should therefore verify those calls are correct. I have chosen to use a test server from  `net/http/httptest` to intercept the requests made by `accountclient` to verify that the calls it makes are correct. This ensures that the tests can run independely of the API that they interact with, preventing them from failing if the API is unavailable.

To help ensure a valid request body is passed to the `create` endpoint I have also included a struct representing Account data. I used this struct to marshal values to JSON to POST to the endpoint. I exposed this struct within the package so that other packages can reuse it for marshalling/unmarshalling