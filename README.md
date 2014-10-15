circlet
=======

circlet is simple support tool for [CircleCI](https://circleci.com/)

Installation
=======


```bash
$ go get github.com/stormcat24/circlet
```

Usage
=======

### Prepare circlet.yml
`circlet.yml` is job definition file.

Top level elements in circlet.yml are `jobs` and `settings`.

Define specific job element under `jobs`. Each job element should have some elements, the follows.

* description: Label of job.
* endpoint: Specify CircleCI's API endpoint. Unnecessary to include querystring.
* method: HTTP method of API (GET/POST/PUT/DELETE)
* query_parameters: Used as querystring. Unnecessary to define `circleci-token` here.
* build_parameters: [Parameterized build parameters](https://circleci.com/docs/parameterized-builds) on CircleCI.

`setting` is global setting of circlet, the follows.

* api_host: Api host of CircleCI
* api_token: Your api token of CircleCI.

##### Example
```Ruby
jobs:
  build_test:
    description: build ${branch}
    endpoint: /project/stormcat24/circlet/tree/${branch}
    method: POST
    build_parameters:
      param1: param1Value
      param2: param2Value
  retry_test:
    description: retry build!!
    endpoint: /project/${repository}/${build_num}/retry
    method: POST
setting:
  api_host: circleci.com
  api_token: WRITE_YOUR_TOKEN_HERE
```

### Execute Job

Execute defined job in circlet.yml by circlet.

##### Command line options

* -c: Path of circlet.yml
* -j: Job name
* -p: Parameters(Pipe-delimited)

##### Example
* execute `build_test`
```bash
$ circlet -c path-to-path/circlet.yml -j build_test -p "branch=master"
```
* execute `retry_test`
```bash
$ circlet -c path-to-path/circlet.yml -j retry_test -p "repository=stormcat/circlet|build_num=5"
```
