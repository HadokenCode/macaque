package rest

import (
	"fmt"
)

const errorsEndpoint = "http://%s"

//Link structure for links
type Link struct {
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}

//Error structure for http errors
type Error struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Link        *Link  `json:"_link,omitempty"`
}

//HTTPError create pointer to Error
func HTTPError(code, description string) *Error {
	return &Error{
		Code:        code,
		Description: description,
		Link: &Link{
			Href: fmt.Sprintf(errorsEndpoint, code),
			Rel:  "self",
		},
	}
}
