// Package main provides a YAML node inspection tool that reads YAML from stdin
// and outputs a detailed analysis of its node structure, including comments
// and content organization.
package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strings"
)

// Constants for YAML node kinds and styles.
const (
	// YAML node kinds as defined in yaml.v3
	kindDocument = int(yaml.DocumentNode)
	kindSequence = int(yaml.SequenceNode)
	kindMapping  = int(yaml.MappingNode)
	kindScalar   = int(yaml.ScalarNode)
	kindAlias    = int(yaml.AliasNode)

	// YAML node styles as defined in yaml.v3
	styleTagged  = int(yaml.TaggedStyle)
	styleDouble  = int(yaml.DoubleQuotedStyle)
	styleSingle  = int(yaml.SingleQuotedStyle)
	styleLiteral = int(yaml.LiteralStyle)
	styleFolded  = int(yaml.FoldedStyle)
	styleFlow    = int(yaml.FlowStyle)

	// indent defines the number of spaces used for indentation
	indent = "  "
)

// main reads YAML from stdin, parses it, and outputs the node structure
func main() {
	reader := io.Reader(os.Stdin)
	dec := yaml.NewDecoder(reader)

	for {
		var node yaml.Node
		err := dec.Decode(&node)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal("Failed to load YAML node:", err)
		}
		fmt.Println(formatNode(node))
	}
}

// formatNode converts a YAML node into a formatted string representation,
// including its kind, value (for scalars), comments, and nested content.
func formatNode(n yaml.Node) string {
	kind := formatKind(int(n.Kind))
	style := formatStyle(int(n.Style))
	tag := formatTag(n.Tag)

	var o string

	o += fmt.Sprintf("Kind: %v\n", kind)

	if style != "" {
		o += fmt.Sprintf("Styl: %v\n", style)
	}
	if n.Anchor != "" {
		o += fmt.Sprintf("Anch: &%v\n", n.Anchor)
	}
	if tag != "" {
		o += fmt.Sprintf("Tag : %v\n", n.Tag)
	}
	if n.HeadComment != "" {
		o += fmt.Sprintf("Head:%v\n", quote(n.HeadComment))
	}
	if n.LineComment != "" {
		o += fmt.Sprintf("Line:%v\n", quote(n.LineComment))
	}
	if n.FootComment != "" {
		o += fmt.Sprintf("Foot:%v\n", quote(n.FootComment))
	}

	if kind == "Scalar" {
		o += fmt.Sprintf("Text:%v\n", quote(n.Value))
	} else {
		o += fmt.Sprintf("More:%v", formatCollection(n.Content))
	}

	return o
}

// quote formats a string as a quoted YAML scalar, preserving newlines.
func quote(s string) string {
	quoted := fmt.Sprintf("\"%v\"", s)
	lines := strings.Split(quoted, "\n")
	if len(lines) == 1 {
		return " " + quoted
	}
	var result string
	for i, line := range lines {
		result += indent + line
		if i < len(lines)-1 {
			result += "\n"
		}
	}
	return "\n" + result
}

// FormatCollection formats a slice of YAML nodes, applying proper indentation
// to maintain the hierarchical structure.
func formatCollection(c []*yaml.Node) string {
	var o string

	for _, node := range c {
		o = fmt.Sprintf(
			"%s\n%s",
			o,
			chomp(
				indentString(
					formatNode(*node),
				),
			),
		)
	}

	return indentString(o)
}

// chomp removes trailing newlines from a string.
func chomp(s string) string {
	return strings.TrimSuffix(s, "\n")
}

// formatKind converts a YAML node kind integer to its string representation.
func formatKind(i int) string {
	switch i {
	case kindDocument:
		return "Document"
	case kindSequence:
		return "Sequence"
	case kindMapping:
		return "Mapping"
	case kindScalar:
		return "Scalar"
	case kindAlias:
		return "Alias"
	default:
		return "Unknown"
	}
}

// formatStyle converts a YAML node style integer to its string representation.
func formatStyle(i int) string {
	switch i {
	case styleTagged:
		return "Tagged"
	case styleDouble:
		return "Double"
	case styleSingle:
		return "Single"
	case styleLiteral:
		return "Literal"
	case styleFolded:
		return "Folded"
	case styleFlow:
		return "Flow"
	}
	return ""
}

// formatTag converts a YAML tag string to its string representation.
func formatTag(tag string) string {
	if tag == "!!str" || tag == "!!map" || tag == "!!seq" {
		return ""
	}
	return tag
}

// indentString adds indentation to each line of the input string.
func indentString(s string) string {
	var result string
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		result += indent + line
		if i < len(lines)-1 || strings.HasSuffix(s, "\n") {
			result += "\n"
		}
	}
	return result
}
