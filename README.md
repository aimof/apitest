# apitest: RESTful APIs testing command. BDD-like, blackbox, automated testing with containers or CIs

* API testing command for containers and CIs.
* Simplest, Fastest, Smallest API test tool.
* BDD like YAML Testing format: Readable and Writable.
* Smallest docker image (It will be about 10MB)
* JSON Format support.
* Mock support(developping)

[![Build Status](https://travis-ci.org/aimof/apitest.svg?branch=master)](https://travis-ci.org/aimof/apitest)

## Purpose: Why I develop apitest?

* I want to use Automatic Sequential API Testing Tool in docker or CI.
* I can't find any docker friendly API Testing tools.
* I don't need GUI, I don't need multifunctional testing.

__So I develop apitest, which is small, docker friendly and minimum.__

## Why and When you use apitest?

* When you develop an API server, try apitest with CI.
* apitest provides your development automatic continuous testing.
* It brings you more safety and faster development.

## apitest vs other API Testing tools or framework.

* Testing Frameworks with GUI: apitest is a minimal simple command for CI.
* Karate: Karate is great tool so apitest is sub-choise when you want to do minimum test in CI.
* Manual Testing: apitest has Repeatablity and doesn't have human errors.

## Example of yaml file

```
Feature: Sample Yaml

Scenarios:
  - Scenario: Normal
    description: Normal
    Given:
      host: http://server:8080
    Tests:
      - When:
          method: GET
          path: /api
        Then:
          status: 200
          format: "empty"
      - When:
          path: "/api/users/"
          method: POST
          body: '{"Token": "Foo", "Name": "Foo"}'
        Then:
          status: 200
          format: application/json
          require:
            - 'match {"name": "Foo", "Info": "#String"}'
      - When:
          path: /api/users/
          method: POST
          body: '{"Token": "Foo"}'
        Then:
          status: 200
          format: application/json
          require:
            - 'match {"users": "#Array"}'
  
```

## DEMO

please run test.sh.

## Install

* using locally
* using with docker

Requirements: Go 1.12.x

```
go get -u github.com/aimof/apitest/cmd/apitest
# workdir: $GOPATH/src/github.com/aimof/cmd/apitest
go install
```

## Usage

```
apitest /path/to/test.yaml
```

### Request body and Response body

* Only Json support.

### Types

```
NumberType = "#Number"
StringType = "#String"
BoolType   = "#Bool"
ObjectType = "#Object"
ArrayType  = "#Array"
NullType   = "#Null"
```

## Testing

```sh
# in the repository root
go test ./...
./test.sh
```

## PLAN

1. Change YAML Format to BDD like Style (Destructive change): Done
2. Make Docker Image: Trial is done.
3. Support Single Mock (Destructive change)
4. Support Multiple Mock (Destructive change)


## Contribution

Please feel free to make Issues or PRs but I will plan some destructive change.

## LICENSE

[MIT](./LICENSE)