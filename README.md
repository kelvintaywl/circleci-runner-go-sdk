# circleci-runner-go-sdk

(Unofficial) Go SDK for CircleCI Runner API

The code here is auto-generated via go-swagger and the _openapi.yaml_ (Swagger 2.0) API definition.

```console
# Assumes you have go-swagger installed outside of this project.
# See https://goswagger.io/install.html#installing-from-binary-distributions
$ swagger generate client -f ./openapi.yaml -A circleci

$ go mod tidy
```

## Validating

```console
# (optional) install swagger-cli
# see https://apitools.dev/swagger-cli/
$ npm install -g @apidevtools/swagger-cli

$ swagger-cli validate ./openapi.yaml
```

## Notes

This is separate from https://github.com/kelvintaywl/circleci-go-sdk.

The reasons being:

1. Runner APIs used here are not public / official on CircleCI Docs yet.
2. (for CircleCI Cloud) Runner APIs are run against https://runner.circleci.com, not https://circleci.com
