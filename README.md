# Rosie Hamilton

I have written the client library in a separate Go package called `accountclient`
This is so that the `accountclient` package can be easily imported and reused by other Go program

The tests for `accountclient` are in package named `accountclient_test` this is so that they can only
interact with exposed interface of `accountclient` i.e. only use public methods (starting with capital letters)

I have placed the code for each client operation `create`, `fetch`, `list` and `delete` in a separate file 
This is to increase readability and aid debugging so if `fetch` caused an error, the dev would know to look in `fetch.go`

I have declared the baseUrl for the API in `main.go` so that if the API version changes, it only needs to be updated once in the code.

I would usually want to mock the HTTP responses from the API for `accountClient` so it's tests can run independently of the API that they interact with. This would prevent the client tests from failing if the API was unavailable. However as this task has been presented inside a docker container, I am making the assumption that this was done because the task requires `accountclient` to interact directly with the API inside Docker. By doing this the tests for `accountclient` are closer to integration tests than unit tests as it is generally considered that unit tests shouldn't change database state.

