# apitest: Sequencial blackbox testing command for APIs.

* Sequencial blackbox automation testing Command for APIs.
* No GUI.
* YAML input: Readable and Writable!
* Automatic sequencial test.

[![Build Status](https://travis-ci.org/aimof/apitest.svg?branch=master)](https://travis-ci.org/aimof/apitest)

## Install

```
go get -u github.com/aimof/apitest/cmd/apitest
go install
```

## YAML Format

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
```

## Request body and Response body

* Basically, you require one request body file and one expected response body file by one request.
* if you don't want to use request body file, `apitest` sends empty body.
* Response pattern file is always needed.

## Response body pattern matching.

* `apitest` only support text format with pattern matching.
* (I'll support json and others soon.)

* Asterisk `*` means wild card in want pattern.
* Asterisk does't restrict the length lf letters.
* (I'll support `?` (one letter matching) soon.)

### example of pattern matching.

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

## Demo & Test

Please clone the repo and run this command in repo root.

```sh
# in the repository root
./demo.sh
```

## Testing

```sh
# in the repository root
go test ./...
```

## ToDo

* [ ] Json Response support.
* [ ] `?` pattern match support.
* [ ] More test cases.
* [ ] Refactor. (especially comment)
* [ ] Timeout.
* [ ] Wildcard character selection.
* [ ] Unittest repo root.

## Contribution

Please feel free to make Issues or PRs.

## LICENSE

MIT