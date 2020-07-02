package providers

// BitbucketPullRequest ...
type BitbucketPullRequest struct {
	Title       string     `json:"title"`      
	Description string     `json:"description"`
	State       string     `json:"state"`      
	Open        bool       `json:"open"`       
	Closed      bool       `json:"closed"`     
	FromRef     Ref        `json:"fromRef"`    
	ToRef       Ref        `json:"toRef"`      
	Locked      bool       `json:"locked"`     
	Reviewers   []Reviewer `json:"reviewers"`  
	Links       Links      `json:"links"`      
}

// Ref ...
type Ref struct {
	ID         string     `json:"id"`        
	Repository Repository `json:"repository"`
}

// Repository ...
type Repository struct {
	Slug    string      `json:"slug"`   
	Name    interface{} `json:"name"`   
	Project Project     `json:"project"`
}

// Project ..
type Project struct {
	Key string `json:"key"`
}

// Links ...
type Links struct {
	Self []interface{} `json:"self"`
}

// Reviewer ...
type Reviewer struct {
	User User `json:"user"`
}

// User ...
type User struct {
	Name string `json:"name"`
}
