// Command openapicheck validates the canonical OpenAPI document.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: openapicheck <openapi-document>")
		os.Exit(2)
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = false
	document, err := loader.LoadFromFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "load OpenAPI document: %v\n", err)
		os.Exit(1)
	}
	if err := document.Validate(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "validate OpenAPI document: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("OpenAPI validation passed: %s\n", os.Args[1])
}
