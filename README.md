# go-enf #

[![Release](https://img.shields.io/github/release/xaptum/go-enf.svg)](https://github.com/xaptum/go-enf/releases)
[![Build Status](https://travis-ci.com/xaptum/go-enf.svg?branch=master)](https://travis-ci.com/xaptum/go-enf)
[![Codecov branch](https://img.shields.io/codecov/c/github/xaptum/go-enf/master.svg)](https://codecov.io/gh/xaptum/go-enf)
[![Go Report Card](https://goreportcard.com/badge/github.com/xaptum/go-enf)](https://goreportcard.com/report/github.com/xaptum/go-enf)

go-enf is a Go client library for accessing the
[Xaptum](https://www.xaptum.com) ENF API.

## Usage ##

```go
import "github.com/xaptum/go-enf/v0/enf" // with go modules enabled (G0111MODULE=on or outside GOPATH)
import "github.com/xaptum/go-enf/enf"    // with go modules disabled
```

Construct a new ENF client, then use the various services on the
client to access different parts of the ENF API. For example:

``` go
const (
    domain = https://demo.xaptum.io
)

// Create a client
client := enf.NewClient(domain, nil)

// List all firewall rules in a network
rules, _, err := client.Firewall.ListRules(context.Background(), "fd00:8f80:0:1::/64")
```

The various services in the client correspond to the structure of the
ENF API documentation and the
[enfcli](https://github.com/xaptum/enfcli) commands.

### Authentication

If set, the `Client.ApiToken` member will be included as the
authorization token for each request. Some ENF API endpoints do not
require authentication, so setting this member is optional.

Use the `Client.Auth` service to request an authorization token using
your username and password.

``` go
const (
    username = "user1"
    password = "password1"
)

// Get an authentication token
authReq := &enf.AuthRequest{Username: &username, Password: &password}
auth, _, _ := client.Auth.Authenticate(context.Background(), authReq)
if err != nil {
    // Handle error
}

client.ApiToken = *auth.Token
```

## Versioning ##

In general, go-enf follows [semver](https://semver.org/) as closely as
possibly for tagging releases.

- Increment the **major version** for incompatible changes to the Go API
  or behavior.
- Increment the **minor version** for backwards-compatible changes to
  the functionality.
- Increment the **patch version** for backwards-compatible bug fixes.

## License ##
Copyright 2019 Xaptum, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not
use this work except in compliance with the License. You may obtain a copy of
the License from the LICENSE.txt file or at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations under
the License.
