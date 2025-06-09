package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var n yaml.Node

	input, _ := io.ReadAll(os.Stdin)
	str := string(input)

	err := yaml.Unmarshal([]byte(str), &n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(FormatNode(n))
}

func FormatNode(n yaml.Node) string {
	kind := FormatKind(int(n.Kind))
	o := fmt.Sprintf("Kind: %v\n", kind)

	if kind == "Scalar" {
		o += fmt.Sprintf("Value: %v\n", n.Value)
	}

	o += fmt.Sprintf(
		`Head Comment: "%v"
Line Comment: "%v"
Foot Comment: "%v"`,
		n.HeadComment,
		n.LineComment,
		n.FootComment,
	)

	if kind != "Scalar" {
		o += fmt.Sprintf(
			"\nContent:%v",
			FormatCollection(n.Content),
		)
	}

	return o
}

func FormatCollection(c []*yaml.Node) string {
	o := ""

	for _, node := range c {
		o = fmt.Sprintf(
			"%s\n%s",
			o,
			indentString(
				FormatNode(*node),
			),
		)
	}

	return indentString(o)
}

func FormatKind(i int) string {
	var kind string
	switch i {
	case 1:
		kind = "Document"
	case 2:
		kind = "Sequence"
	case 4:
		kind = "Mapping"
	case 8:
		kind = "Scalar"
	case 16:
		kind = "Alias"
	}
	return kind
}

func indentString(s string) string {
	indentStr := "  "
	var result string
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		result += indentStr + line + "\n"
	}
	return result
}
