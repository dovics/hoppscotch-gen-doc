package generator

import (
	"encoding/json"
	"fmt"
	"strings"
)

// HoppscotchCollection represents the root structure of a Hoppscotch collection
type HoppscotchCollection struct {
	Version    int                       `json:"v"`
	Name       string                    `json:"name"`
	Folders    []Folder                  `json:"folders"`
	Requests   []Request                 `json:"requests"`
	Auth       Auth                      `json:"auth"`
	Headers    []Header                  `json:"headers"`
	Variables  []Variable                `json:"variables"`
	Desc       string                    `json:"description"`
}

type Folder struct {
	Version    int       `json:"v"`
	Name       string    `json:"name"`
	Folders    []Folder  `json:"folders"`
	Requests   []Request `json:"requests"`
	Auth       Auth      `json:"auth"`
	Headers    []Header  `json:"headers"`
	Variables  []Variable `json:"variables"`
	Desc       string    `json:"description"`
}

type Request struct {
	Version           string              `json:"v"`
	Name              string              `json:"name"`
	Method            string              `json:"method"`
	Endpoint          string              `json:"endpoint"`
	Params            []Param             `json:"params"`
	Headers           []Header            `json:"headers"`
	PreRequestScript  string              `json:"preRequestScript"`
	TestScript        string              `json:"testScript"`
	Auth              Auth                `json:"auth"`
	Body              Body                `json:"body"`
	RequestVariables  []Variable          `json:"requestVariables"`
	Responses         map[string]Response `json:"responses"`
	Description       string              `json:"description"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type Body struct {
	ContentType string `json:"contentType"`
	Body        string `json:"body"`
}

type Auth struct {
	AuthType   string `json:"authType"`
	AuthActive bool   `json:"authActive"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type Response struct {
	Name string `json:"name"`
}

// MarkdownGenerator generates markdown documentation from Hoppscotch collection
type MarkdownGenerator struct {
	collection HoppscotchCollection
	builder    strings.Builder
}

// Generate parses JSON data and returns markdown documentation
func Generate(data []byte) (string, error) {
	var collection HoppscotchCollection
	if err := json.Unmarshal(data, &collection); err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}

	gen := &MarkdownGenerator{
		collection: collection,
		builder:    strings.Builder{},
	}

	return gen.generate(), nil
}

func (g *MarkdownGenerator) generate() string {
	g.writeTitle()
	g.writeDescription()
	g.writeTOC()
	g.writeContent()
	return g.builder.String()
}

func (g *MarkdownGenerator) writeTitle() {
	fmt.Fprintf(&g.builder, "# %s\n\n", g.collection.Name)
}

func (g *MarkdownGenerator) writeDescription() {
	if g.collection.Desc != "" {
		fmt.Fprintf(&g.builder, "## Description\n\n%s\n\n", g.collection.Desc)
	}
}

func (g *MarkdownGenerator) writeTOC() {
	g.builder.WriteString("## Table of Contents\n\n")

	// Write folders in TOC
	if len(g.collection.Folders) > 0 {
		g.writeFoldersTOC(g.collection.Folders, 1)
	}

	// Write root level requests in TOC
	if len(g.collection.Requests) > 0 {
		g.builder.WriteString("### General\n\n")
		for _, req := range g.collection.Requests {
			anchor := g.createAnchor(req.Name)
			fmt.Fprintf(&g.builder, "- [%s](#%s)\n", req.Name, anchor)
		}
		g.builder.WriteString("\n")
	}
}

func (g *MarkdownGenerator) writeFoldersTOC(folders []Folder, level int) {
	indent := strings.Repeat("  ", level-1)
	for _, folder := range folders {
		folderAnchor := g.createAnchor(folder.Name)
		fmt.Fprintf(&g.builder, "%s- [%s](#%s)\n", indent, folder.Name, folderAnchor)

		// Write requests in this folder
		if len(folder.Requests) > 0 {
			reqIndent := strings.Repeat("  ", level)
			for _, req := range folder.Requests {
				reqAnchor := g.createAnchor(req.Name)
				fmt.Fprintf(&g.builder, "%s- [%s](#%s)\n", reqIndent, req.Name, reqAnchor)
			}
		}

		// Recursively write subfolders
		if len(folder.Folders) > 0 {
			g.writeFoldersTOC(folder.Folders, level+1)
		}
	}
	g.builder.WriteString("\n")
}

func (g *MarkdownGenerator) writeContent() {
	// Write folders with their requests
	if len(g.collection.Folders) > 0 {
		g.writeFolders(g.collection.Folders, 2)
	}

	// Write root level requests
	if len(g.collection.Requests) > 0 {
		g.builder.WriteString("## General\n\n")
		for _, req := range g.collection.Requests {
			g.writeRequest(req)
			g.builder.WriteString("\n---\n\n")
		}
	}
}

func (g *MarkdownGenerator) writeFolders(folders []Folder, level int) {
	for _, folder := range folders {
		// Write folder heading
		headingPrefix := strings.Repeat("#", level)
		fmt.Fprintf(&g.builder, "%s %s\n\n", headingPrefix, folder.Name)

		// Write folder description
		if folder.Desc != "" {
			fmt.Fprintf(&g.builder, "%s\n\n", folder.Desc)
		}

		// Write requests in this folder
		for _, req := range folder.Requests {
			g.writeRequest(req)
			g.builder.WriteString("\n---\n\n")
		}

		// Recursively write subfolders
		if len(folder.Folders) > 0 {
			g.writeFolders(folder.Folders, level+1)
		}
	}
}

func (g *MarkdownGenerator) writeRequest(req Request) {
	// Request name as heading
	fmt.Fprintf(&g.builder, "### %s\n\n", req.Name)

	// Description
	if req.Description != "" {
		fmt.Fprintf(&g.builder, "**Description:**\n\n%s\n\n", req.Description)
	}

	// HTTP Method and Endpoint
	methodBadge := g.getMethodBadge(req.Method)
	fmt.Fprintf(&g.builder, "%s\n\n", methodBadge)
	fmt.Fprintf(&g.builder, "**Endpoint:** `%s`\n\n", req.Endpoint)

	// Headers
	if len(req.Headers) > 0 {
		g.builder.WriteString("#### Headers\n\n")
		g.builder.WriteString("| Key | Value | Description |\n")
		g.builder.WriteString("|-----|-------|-------------|\n")
		for _, h := range req.Headers {
			desc := h.Desc
			if desc == "" {
				desc = "-"
			}
			fmt.Fprintf(&g.builder, "| %s | %s | %s |\n", h.Key, h.Value, desc)
		}
		g.builder.WriteString("\n")
	}

	// Query Parameters
	if len(req.Params) > 0 {
		g.builder.WriteString("#### Query Parameters\n\n")
		g.builder.WriteString("| Key | Value | Description |\n")
		g.builder.WriteString("|-----|-------|-------------|\n")
		for _, p := range req.Params {
			desc := p.Desc
			if desc == "" {
				desc = "-"
			}
			fmt.Fprintf(&g.builder, "| %s | %s | %s |\n", p.Key, p.Value, desc)
		}
		g.builder.WriteString("\n")
	}

	// Request Body
	if req.Body.ContentType != "" && req.Body.Body != "" {
		g.builder.WriteString("#### Request Body\n\n")
		fmt.Fprintf(&g.builder, "**Content-Type:** %s\n\n", req.Body.ContentType)
		g.builder.WriteString("```json\n")
		// Try to format JSON if possible
		var formattedJSON interface{}
		if err := json.Unmarshal([]byte(req.Body.Body), &formattedJSON); err == nil {
			if pretty, err := json.MarshalIndent(formattedJSON, "", "  "); err == nil {
				g.builder.WriteString(string(pretty))
			} else {
				g.builder.WriteString(req.Body.Body)
			}
		} else {
			g.builder.WriteString(req.Body.Body)
		}
		g.builder.WriteString("\n```\n\n")
	}

	// Authentication
	if req.Auth.AuthType != "" && req.Auth.AuthType != "inherit" && req.Auth.AuthType != "none" {
		g.builder.WriteString("#### Authentication\n\n")
		fmt.Fprintf(&g.builder, "**Type:** %s\n\n", req.Auth.AuthType)
	}
}

func (g *MarkdownGenerator) getMethodBadge(method string) string {
	colors := map[string]string{
		"GET":     "🟢",
		"POST":    "🟡",
		"PUT":     "🔵",
		"PATCH":   "🟠",
		"DELETE":  "🔴",
		"HEAD":    "⚪",
		"OPTIONS": "⚫",
	}

	emoji := colors[method]
	if emoji == "" {
		emoji = "⚪"
	}

	return fmt.Sprintf("**%s %s**", emoji, method)
}

func (g *MarkdownGenerator) createAnchor(text string) string {
	anchor := strings.ToLower(text)
	anchor = strings.ReplaceAll(anchor, " ", "-")
	anchor = strings.ReplaceAll(anchor, "/", "-")
	anchor = strings.ReplaceAll(anchor, "_", "-")
	// Remove special characters
	var result strings.Builder
	for _, r := range anchor {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			result.WriteRune(r)
		}
	}
	return result.String()
}
