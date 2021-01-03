// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

/**
 * 有关 JSON 的 Github 接口练习
 *
 * @author moqi
 * On 2021/1/3 20:50
 */
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
