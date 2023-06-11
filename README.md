# Jobadder API Utility Written In Go

[![Test](https://github.com/xehapa/jago/actions/workflows/pull_request.yml/badge.svg?event=push)](https://github.com/xehapa/jago/actions/workflows/pull_request.yml)

--- Folder Structure ---
```
LICENSE
README.md
.env.test
.env.dist
[api]
    ├── auth.go
    └── jobadder.go
[config]
    └── config.go
go.mod
go.sum
main.go
[models]
    ├── auth.go
    └── job.go
[runner]
    └── runner.go
[tests]
    └── [unit]
        ├── config_test.go
        ├── httpclient_test.go
        ├── jobadder_test.go
        └── main_test.go
[utils]
    └── httpclient.go
```