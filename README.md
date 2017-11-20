# go-coverage-example

Code example to go coverage tools.
* go test -cover : basic test coverage tool.
* go test -coverprofile=coverage.out : create coverage file used to present covered lines in html page.
* go tool cover -html=coverage.out : view html page with details about covered lines. 