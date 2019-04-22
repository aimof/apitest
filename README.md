# apitest: RESTful APIs testing command. BDD like, blackbox, automated testing with containers or CIs

* API testing command for containers and CIs.
* Simplest, Fastest, Smallest API test tool.
* BDD like YAML Testing format: Readable and Writable.
* Smallest docker image (It will be about 10MB)

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

## DEMO

```
$ apitest ./test.yaml
2019/04/07 10:12:38 Test case 0 : Success
2019/04/07 10:12:38 Test case 1 : Success
2019/04/07 10:12:38 Test case 2 : Success
2019/04/07 10:12:38 Test case 3 : Success
2019/04/07 10:12:38 Test case 4 : Success
2019/04/07 10:12:38 Test case 5 : Success
2019/04/07 10:12:38 Test case 5 : Fail  Plaese Read Log.
2019/04/07 10:12:38 Stopped: Test case 5 Failed
```

If you want to try it, please clone the repo and run this command in repo root.

```sh
# in the repository root
./demo.sh
```

Sample of docker-compose.yml is [here](./demo/docker-compose.yml)

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

### YAML Format

```yaml
# config does not work now.
config:

# testcases (array of object): one object has one http request.
testcases:
  - Name: case0               # testcase name.
    URL: http://localhost/api # URL you want to kick.
    method: POST              # http method.
    header:                   # header (array of string): separated by collon.
      - Key:Value0:Value1
    bodypath: ./body.txt      # path to request body file.
    want:                     # want is response you want.
      statuscode: 200         # statuscode you expected. (int)
      bodypath: ./pattern.txt # path to expected response pattern file. Read below.
  - Name: case1
    URL: http://localhost/api
  ... # Repeat this format.
```

### Request body and Response body

* Basically, you require one request body file and one expected response body file by one request.
* if you don't want to use request body file, `apitest` sends empty body.
* Response pattern file is always needed.

### Response body pattern matching.

* `apitest` only support text format with pattern matching.
* (I'll support json and others soon.)

* Asterisk `*` means wild card in want pattern.
* Asterisk does't restrict the length lf letters.
* (I'll support `?` (one letter matching) soon.)

#### example of pattern matching.

```
# match
got:  "FizzBuzz"
want: "FizzBuzz"

# match
got:  "FizzBuzz"
want: "F*z"

# match
got:  "FizzBuzz"
want: "Fizz*"

# match
got:  "FizzBuzz"
want: "*"

# not match
got:  "FizzBuzz"
want: "fizzbuzz:

# not match
got:  "FizzBuzz"
want: "F*zz"
```

## Testing

```sh
# in the repository root
go test ./...
```

## PLAN

1. Change YAML Format to BDD like Style (Destructive change): In progress
2. Make Docker Image: Trial is done.
3. Support Single Mock (Destructive change)
4. Support Multiple Mock (Destructive change)


## Contribution

Please feel free to make Issues or PRs.

## LICENSE

[MIT](./LICENSE)