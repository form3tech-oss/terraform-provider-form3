  [![Build Status](https://travis-ci.org/ewilde/go-form3.svg?branch=master)](https://travis-ci.org/ewilde/go-form3)
# Developing
## Build locally
* `git clone git@github.com:ewilde/go-form3.git`
* `cd go-form3`
* `make build`

## Running tests
# Tests

| Environment variables| Description                                |
|:---------------------|:-------------------------------------------|
| FORM3_HOST           | Form 3 host e.g. api.form3.tech            |
| FORM3_ACC            | Set to `1` to run integration tests        |
| FORM3_CLIENT_ID      | Client id                                  |
| FORM3_CLIENT_SECRET  | Secret                                     |
| TF_LOG               | "TRACE", "DEBUG", "INFO", "WARN", "ERROR"  |

## Example
`make test FORM3_ACC=1 FORM3_CLIENT_ID=xxx FORM3_CLIENT_SECRET=xxx FORM3_HOST=api.whistle.env.form3.tech FORM3_ORGANISATION_ID=xxx`


# Swagger
To generate the swagger model run: `swagger generate client -f ./swagger.yaml`
