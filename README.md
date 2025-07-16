# SnapshotGo

A minimal **Go SDK** for interacting with the **[Snapshot](https://snapshot.org)** off-chain governance Hub ([GraphQL API](https://hub.snapshot.org/graphql)).\
SnapshotGo handles request throttling, exposes a simple `Query` helper, and plays nicely with Go’s `context` and generics so you can fetch **spaces, proposals, votes, follows, roles, messages, voting power**—or anything else the API supports—using plain Go structs.

---

## Table of Contents
1. [Features](#features)  
2. [Installation](#installation)  
3. [Quick Start](#quick-start)  
4. [Usage](#usage)  
   - [Custom GraphQL Queries](#custom-graphql-queries)  
   - [Rate Limiting](#rate-limiting)  
5. [Examples](#examples)  
6. [Roadmap](#roadmap)  
7. [Contributing](#contributing)  
8. [License](#license)  

---

## Features
- **Tiny API** – one constructor + one generic `Query` function.  
- **Built-in rate limiter** (defaults to the 60 req/min public limit; overridable).  
- **Idiomatic Go** – explicit context handling, timeouts, strongly typed responses.  
- **Generic responses** – supply *any* struct and have the JSON unmarshalled for you.  
- **Production-ready HTTP client** – sane timeouts, easy to swap for your own.  

> **Why?**  
> Snapshot’s official JavaScript client is great for web apps, but Go back-ends, cloud functions, and CLIs need a lightweight alternative—without pulling in an entire JS runtime.

---

## Installation
```bash
go get github.com/This-Is-Prince/snapshotgo
```

> Go ≥ 1.21 is recommended (for improved generics & `context` features).

---

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/This-Is-Prince/snapshotgo"
)

const PROPOSAL_QUERY = `
query Proposal($id: String!) {
  proposal(
    id: $id
  ) {
    id
    title
    body
    choices
    start
    end
    snapshot
    state
    author
    space {
      id
      name
    }
  }
}
`

const PROPOSAL_ID = "0x586de5bf366820c4369c041b0bbad2254d78fafe1dcc1528c1ed661bb4dfb671"

type ProposalResponse struct {
	Proposal Proposal `json:"proposal"`
}

type Proposal struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	snapshotHub := snapshotgo.NewSnapshot()

	args := snapshotgo.Args{
		Query: PROPOSAL_QUERY,
		Variables: map[string]interface{}{
			"id": PROPOSAL_ID,
		},
	}

	var data ProposalResponse

	err := snapshotgo.Query(snapshotHub, args, &data)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	fmt.Printf("%#v", data)
}

```

---

## Usage

### Custom GraphQL Queries

1. Craft a query in the [GraphiQL Explorer](https://hub.snapshot.org/graphql).
2. Copy it into your Go code (use raw-string literals `` `like this` ``).
3. Define a Go struct mirroring the shape of the expected JSON.
4. Call `snapshotgo.Query(...)` and pass a pointer to your struct.

The SDK never imposes a fixed schema—**you control** the query and the response type.

### Rate Limiting

| Setting        | Default                 | Purpose                                                                                                      |
| -------------- | ----------------------- | ------------------------------------------------------------------------------------------------------------ |
| `InitialBurst` | `1` request             | *Bucket depth* – how many calls can fire instantly.                                                          |
| `RateLimit`    | `1 request / 2 seconds` | *Steady-state rate* (≈ 30 req/min).<br/>Change to `rate.Every(time.Second)` for the public 60 req/min limit. |
| `IsLimited`    | `true`                  | Toggle throttling entirely (e.g., when using an API key with higher limits).                                 |

```go
snap := snapshotgo.NewSnapshot()
snap.InitialBurst = 5
snap.RateLimit    = rate.Every(time.Second) // 60 rpm
snap.Limiter      = rate.NewLimiter(snap.RateLimit, snap.InitialBurst)
```

---

## Examples

| Example          | File                              | Description                                      |
| ---------------- | --------------------------------- | ------------------------------------------------ |
| Fetch a proposal | `examples/proposal/main.go`       | Retrieves a single proposal by ID.        |
| Fetch spaces   | `examples/spaces/main.go` | Retrieves spaces.                  |
| Custom client    | `examples/custom_http/main.go`    | Swap in a custom `http.Client` |

Clone the repo and run any example:

```bash
go run ./examples/proposal
```

---

## Roadmap

* [ ] **Automatic pagination helpers** (`VotesIterator`, `SpacesIterator`)
* [ ] **Typed wrappers** for common entities (spaces, proposals, votes).
* [ ] **Context cancellation hooks** for graceful shutdown in long-running apps.
* [ ] **CI pipeline** (lint, staticcheck, test coverage).
* [ ] **GitHub Actions** to publish a versioned `go` module.

---

## Contributing

Pull requests, issues, and discussions are very welcome!

1. Fork the repository.
2. `git checkout -b feat/my-awesome-feature`
3. Commit your changes (`git commit -am "Add 🤯 feature"`).
4. Push to the branch (`git push origin feat/my-awesome-feature`).
5. Open a PR and describe **why** & **how**.

**Dev scripts**

```bash
go vet ./...
go test ./...
golangci-lint run
```

---

## License

SnapshotGo is released under the **MIT License** – see [LICENSE](LICENSE) for details.

---

### Star 🌟 the repo if you find it useful!
