

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

// GitHubAPIURL is  URL of API
const GitHubAPIURL = "https://api.github.com/repos/"

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

// NewIssue has minimum field of GitHubIssue
type NewIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CloseIssue has issue status.
type CloseIssue struct {
	State string `json:"state"`
}
