# Package `naming`

[![Go Reference](https://pkg.go.dev/badge/github.com/instantup/naming.svg)](https://pkg.go.dev/github.com/instantup/naming)

---

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
