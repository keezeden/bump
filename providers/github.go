package providers

// GithubPullRequest ...
type GithubPullRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"` 
	Head  string `json:"head"` 
	Base  string `json:"base"` 
}
