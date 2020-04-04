## Introduction
A [Golang](https://golang.org/)-based tool to convert OpenAPI spec into `.docx` based format.

Licensed under [MIT License](LICENSE)

## Features
1. Convert OpenAPI spec into .docx format
2. Support OpenAPI 3.0 version
3. Feature to automatically push the result into Google Drive*1

*1 need to have Google Cloud Platform account with additional steps required. Tutorial below.

## Pre-requisites
- Go >= v1.14

## How-to-use
1. Make sure you have fulfilled the requirements above
2. Git clone `https://github.com/weisurya/go-openapi-converter`
3. Go build `go build -i -o go-openapi-converter app/main.go`
4. Type `go-openapi-converter -h` to learn more about the available commands
5. Type `convert -s sample.v1.yaml -t template/standard.docx -o result.sample.docx` to convert OpenAPI spec into .docx format
6. Type `upload -c credentials.sample.json -d directory -e foo@gmail.com -f result.sample.docx` to upload file into Google Drive
7. To learn more about each command, please type `-h`


## Tutorial - Push result to Google Drive
1. Use Google Cloud Platform. If you don't have any account, please create it first.
2. Go to console > create new project
3. On menu, choose `APIs & Services` > choose `Library`
4. Search `Google Drive API` > click `Enable`
5. If this is your first time, you need to create OAuth Consent Screen. Go to `APIs & Services` > choose `OAuth consent screen`
6. Create new app > set user type as `External`, set the scope for `Google APIs` based on your needs > click `Save`
7. Back to `APIs & Services` > choose `Credentials` > click `+ Create Credentials` > choose `Service account`
8. Fill up the `Service account detail` based on your preference > set permission based on your preference (i.e. Project Owner) > Create key in JSON format > store the result


## Known Limitation
1. Only support for `application/json` based schema
2. Not support nested `$ref`
3. The order of endpoint between table of content and list of endpoint hasn't same yet
4. Numbering list does not appear on Google Docs


## Main Dependencies
- [cobra](https://pkg.go.dev/github.com/spf13/cobra)
- [Google Drive v3](https://pkg.go.dev/google.golang.org/api/drive/v3)
- [kin-openapi](https://pkg.go.dev/github.com/getkin/kin-openapi/openapi3)
- [unioffice](https://github.com/unidoc/unioffice)