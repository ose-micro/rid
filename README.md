# ose-rid

> A UUID-length readable ID generator for your products and microservices.

[![Go Reference](https://pkg.go.dev/badge/github.com/ose-micro/core.svg)](https://pkg.go.dev/github.com/ose-micro/rid)
[![Go Report Card](https://goreportcard.com/badge/github.com/ose-micro/core)](https://goreportcard.com/report/github.com/ose-micro/rid)
[![License](https://img.shields.io/github/license/ose-micro/rid)](LICENSE)

## Features
- Generates **36-character IDs**.
- Includes **prefix**, **timestamp**, and **random uniqueness**.
- Thread-safe and easy to use.
- MIT licensed for free commercial and private use.

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/ose-micro/rid"
)

func main() {
	gen := rid.New("PRD", true)
	fmt.Println(gen.String())
	// Example output: PRD250827-1423-ABCD-EF01-1234567890AB
}
```

## 📦 Install

```bash
go get github.com/ose-micro/rid
```
