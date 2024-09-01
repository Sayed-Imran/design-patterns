package models

type APIUser struct {
    UserName  string `json:"username,omitempty"`
    FirstName string `json:"first_name,omitempty"`
    LastName  string `json:"last_name,omitempty"`
    Email     string `json:"email,omitempty"`
}
