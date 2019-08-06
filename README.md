# gody
A lightweight body manipulator for Go

## Quick start

```go
package main

import (
	"log"

	"github.com/guiferpa/gody"
	"github.com/guiferpa/gody/rule"
)

type Body struct {
	Text string `json:"text" validate:"required=true"`
}

func main() {
	b := Body{}

	valid, err := gody.Validate(b, nil)
	if err != nil {
		if !valid {
			log.Println("body do not validated", err)
			return
		}

		switch err.(type) {
		case *rule.ErrRequired:
			log.Println("required error", err)
			break
		}
	}
}
```

> You can access more [examples](https://github.com/guiferpa/gody/blob/master/example/validate.go)
