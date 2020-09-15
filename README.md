# jimdn/objectid
A package to generate a 12-bytes unique object identifier.

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimdn/objectid?style=flat-square)](https://goreportcard.com/report/github.com/jimdn/objectid)

## Installation

Use go get.

	go get github.com/jimdn/objectid

Then import the package into your own code.

	import "github.com/jimdn/objectid"


## Example

```go
package main

import (
	"fmt"
	"github.com/jimdn/objectid"
)

func main() {
	objIdStr := objectid.New().String()
	fmt.Printf("%s\n", objIdStr)

	objId, _ = objectid.Parse(objIdStr)
	fmt.Printf("%d-%d-%d-%d\n", objId.Timestamp(), objId.Machine(), objId.Pid(), objId.Increment())
}
```
