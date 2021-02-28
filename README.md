# go-mpos
 [![GoDoc reference](https://img.shields.io/badge/docs-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/thelolagemann/go-mpos)  [![GitHub](https://img.shields.io/github/license/thelolagemann/go-mpos?style=flat-square)](https://github.com/thelolagemann/gompress/blob/main/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thelolagemann/go-mpos?style=flat-square&l)

> An unofficial [PHP-MPOS](https://github.com/MPOS/php-mpos) API Client for go.

## Getting started

### Install
`go get -u github.com/thelolagemann/go-mpos`

### Example usage
```golang
import "github.com/thelolagemann/go-mpos"

pool := NewMiningPoolHub("API_KEY", "BASE_URL")
workers, err := pool.UserWorkers()
if err != nil {
    // handle err
}
for _, worker := range workers {
    fmt.Printf("%v is hashing away at %v Mh/s", worker.Username, worker.Hashrate)
}
```
