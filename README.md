# go-imgur
[![GoDoc](https://godoc.org/github.com/koffeinsource/go-imgur?status.svg)](https://godoc.org/github.com/koffeinsource/go-imgur)
[![Build Status](https://travis-ci.org/koffeinsource/go-imgur.svg?branch=master)](https://travis-ci.org/koffeinsource/go-imgur)
[![Go Report Card](https://goreportcard.com/badge/github.com/koffeinsource/go-imgur)](https://goreportcard.com/report/github.com/koffeinsource/go-imgur)
[![Coverage Status](
https://coveralls.io/repos/github/koffeinsource/go-imgur/badge.svg?branch=master)](https://coveralls.io/github/koffeinsource/go-imgur?branch=master)

Go library to use the imgur.com API. At the moment only the anonymous part of the API is supported, but that is used in a production environment.

### Versions

We use gopkg.in for versioning, but are still at `Version 0` and do not offer any API stability guarantees. We are not aware of any major API issues and may switch to `Version 1` soon. If you have ideas or feature requests please let us know!

Please use the following import:
```go
import "gopkg.in/koffeinsource/go-imgur.v0"
```

### Example

To see some simple example code please take a look at the command line client found in `imgurcmd/main.go`.
