
// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

import "time"

// IssuesURL URL of API
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult search result
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue issue information
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User user information
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
