# Package `naming`

[![GitHub Workflow Status (master)](https://img.shields.io/github/workflow/status/instantup/naming/Build/master?style=flat-square)](https://github.com/instantup/naming/actions?query=workflow:Build+branch:master) [![Go Report Card](https://goreportcard.com/badge/github.com/instantup/naming?style=flat-square)](https://goreportcard.com/report/github.com/instantup/naming) [![Go Reference](https://pkg.go.dev/badge/github.com/instantup/naming)](https://pkg.go.dev/github.com/instantup/naming)

## Installation

```shell
go get github.com/instantup/naming
```

## Overview

Package `naming` provides functions for name splitting and converting to common naming conventions.

### Example (Naming Conventions)

Code:

```go
package main

import (
	"fmt"
	"github.com/instantup/naming"
)

func main() {
	for _, formatted := range []string{
		naming.Flat("Alice-WasBeginning"),
		naming.Upper("toGetVery"),
		naming.Mixed("tired-OF sitting"),
		naming.UpperMixed("by__herSister"),
		naming.Snake("_ONThe bank,"),
		naming.UpperSnake("andOfHaving"),
		naming.Kebab("nothingTo do:"),
	} {
		fmt.Println(formatted)
	}
}
```

Output:

```
alicewasbeginning
TOGETVERY
tiredOfSitting
ByHerSister
on_the_bank
AND_OF_HAVING
nothing-to-do
```
