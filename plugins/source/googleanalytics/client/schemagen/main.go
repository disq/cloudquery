package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/client"
	"github.com/cloudquery/codegen/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	jsonschema.GenerateIntoFile(new(client.Spec), path.Join(currDir(), "..", "schema.json"))
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
