# GO SDK for [blrec](https://github.com/acgnhiki/blrec)

## Update API Specification
1. Alter `openapi.json` file to new version (http://localhost:2233/openapi.json)
2. Run command:
```shell
go generate
```
3. Check `client.gen.go` file.

## Usage
### Add as dependency
```shell
go get github.com/wintbiit/blrec-sdk-go
```
```go
package main

import (
    "fmt"
    brec "github.com/wintbiit/blrec-sdk-go"
)

func main() {
	client, _ := brec.NewBlrecClient("http://localhost:2356");
	// or client, _ := brec.NewBrecClientWithAuth("http://localhost:2356", "<your_api_Key>");
}
```