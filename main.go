package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"io/ioutil"
	"log"
)

func main() {
	tfFile := "./main.hcl"
	fileContent, err := ioutil.ReadFile(tfFile)
	if err != nil {
		log.Fatal("Failed to read file: %v", err)

	}
	parser := hclparse.NewParser()

	parsedHCL, diagnostics := parser.ParseHCL(fileContent, tfFile)
	if diagnostics.HasErrors() {
		log.Fatalf("Failed to parse HCL: %v", diagnostics)
	}

	var variables = hcldec.Variables(parsedHCL.Body, tfFile)

	for _, variable := range variables {
		fmt.Printf("Variable: %s, Type: %s, Default: %s\n", variable.Name(), variable.Type(), variable.Default)
	}
}
