# go-zendesk
[![Gitter](https://badges.gitter.im/terraform-provider-zendesk/Lobby.svg)](https://gitter.im/terraform-provider-zendesk/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Build Status](https://travis-ci.org/fairyhunter13/go-zendesk.svg?branch=master)](https://travis-ci.org/fairyhunter13/go-zendesk)
[![Build status](https://ci.appveyor.com/api/projects/status/ce4p1mswjkdftv6o/branch/master?svg=true)](https://ci.appveyor.com/project/fairyhunter13/go-zendesk/branch/master)
[![Coverage Status](https://coveralls.io/repos/github/fairyhunter13/go-zendesk/badge.svg?branch=master)](https://coveralls.io/github/fairyhunter13/go-zendesk?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fairyhunter13/go-zendesk)](https://goreportcard.com/report/github.com/fairyhunter13/go-zendesk)
[![GoDoc](https://godoc.org/github.com/zenform/go-zendesk?status.svg)](https://godoc.org/github.com/zenform/go-zendesk)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffairyhunter13%2Fgo-zendesk.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffairyhunter13%2Fgo-zendesk?ref=badge_shield)

Zendesk API client library for Go

- [Reference](https://godoc.org/github.com/fairyhunter13/go-zendesk)

## Installation

``` shell
$ go get github.com/fairyhunter13/go-zendesk
```

## Usage

```go
package main

import (
    "context"

    "github.com/fairyhunter13/go-zendesk/zendesk"
)

func main() {
    // You can set custom *http.Client here
    client := zendesk.NewClient(nil)

    // example.zendesk.com
    client.SetSubdomain("example")

    // Authenticate with API token
    client.SetCredential(zendesk.NewAPITokenCredential("john.doe@example.com", "apitoken"))

    // Authenticate with agent password
    client.SetCredential(zendesk.NewBasicAuthCredential("john.doe@example.com", "password"))

    // Create resource
    client.CreateGroup(context.Background(), zendesk.Group{
        Name: "support team",
    })
}
```

## Authors
- [fairyhunter13](https://github.com/fairyhunter13)
- [tamccall](https://github.com/tamccall)

## License

MIT License.

See the file [LICENSE](./LICENSE).


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffairyhunter13%2Fgo-zendesk.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffairyhunter13%2Fgo-zendesk?ref=badge_large)
