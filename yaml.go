// Package main provides a YAML node inspection tool that reads YAML from stdin
// and outputs a detailed analysis of its node structure, including comments
// and content organization.
package main

import (
	"errors"
	"fmt"
	"go.yaml.in/yaml/v3"
	"io"
	"log"
	"os"
)

// NodeInfo represents the information about a YAML node
type NodeInfo struct {
	Kind    string      `yaml:"kind"`
	Style   string      `yaml:"style,omitempty"`
	Anchor  string      `yaml:"anchor,omitempty"`
	Tag     string      `yaml:"tag,omitempty"`
	Head    string      `yaml:"head,omitempty"`
	Line    string      `yaml:"line,omitempty"`
	Foot    string      `yaml:"foot,omitempty"`
	Text    string      `yaml:"text,omitempty"`
	Content []*NodeInfo `yaml:"content,omitempty"`
}

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
		info := formatNode(node)
		output, err := yaml.Marshal(info)
		if err != nil {
			log.Fatal("Failed to marshal node info:", err)
		}
		fmt.Print(string(output))
	}
}

// formatNode converts a YAML node into a NodeInfo structure
func formatNode(n yaml.Node) *NodeInfo {
	info := &NodeInfo{
		Kind: formatKind(n.Kind),
	}

	if style := formatStyle(n.Style); style != "" {
		info.Style = style
	}
	if n.Anchor != "" {
		info.Anchor = n.Anchor
	}
	if tag := formatTag(n.Tag); tag != "" {
		info.Tag = tag
	}
	if n.HeadComment != "" {
		info.Head = n.HeadComment
	}
	if n.LineComment != "" {
		info.Line = n.LineComment
	}
	if n.FootComment != "" {
		info.Foot = n.FootComment
	}

	if info.Kind == "Scalar" {
		info.Text = n.Value
	} else if n.Content != nil {
		info.Content = make([]*NodeInfo, len(n.Content))
		for i, node := range n.Content {
			info.Content[i] = formatNode(*node)
		}
	}

	return info
}

// formatKind converts a YAML node kind into its string representation.
func formatKind(k yaml.Kind) string {
	switch k {
	case yaml.DocumentNode:
		return "Document"
	case yaml.SequenceNode:
		return "Sequence"
	case yaml.MappingNode:
		return "Mapping"
	case yaml.ScalarNode:
		return "Scalar"
	case yaml.AliasNode:
		return "Alias"
	default:
		return "Unknown"
	}
}

// formatStyle converts a YAML node style into its string representation.
func formatStyle(s yaml.Style) string {
	switch s {
	case yaml.TaggedStyle:
		return "Tagged"
	case yaml.DoubleQuotedStyle:
		return "Double"
	case yaml.SingleQuotedStyle:
		return "Single"
	case yaml.LiteralStyle:
		return "Literal"
	case yaml.FoldedStyle:
		return "Folded"
	case yaml.FlowStyle:
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
