  [![Build Status](https://travis-ci.org/ewilde/go-form3.svg?branch=master)](https://travis-ci.org/ewilde/go-form3)
# Developing
## Build locally
* `git clone git@github.com:ewilde/go-form3.git`
* `cd go-form3`
* `make build`

## Running tests
To run the integration test you need to export 3 environment variables
`env FORM3_CLIENT_ID=xxx FORM3_CLIENT_SECRET=xxx FORM3_HOST=api.tabla.env.form3.tech FORM3_ORGANISATION_ID=xxx make`


# Swagger
To generate the swagger model run: `go-swagger generate client -f ./swagger.yaml`
