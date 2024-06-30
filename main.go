package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclparse"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Print("-------------------\n")

	tfFile := "./main.hcl"
	fileContent, err := os.ReadFile(tfFile)
	check(err)
	fmt.Print(string(tfFile))

	parser := hclparse.NewParser()

	parsedHCL, diagnostics := parser.ParseHCL(fileContent, tfFile)
	if diagnostics.HasErrors() {
		log.Fatalf("Failed to parse HCL: %v", diagnostics)
	}

	fmt.Print(parsedHCL)
	// ./main.hcl&{map[] [0xc0000ea000] map[] map[] ./main.hcl:1,1-4,2 ./main.hcl:4,2-2}
	fmt.Print("\n-------------------\n")
}
