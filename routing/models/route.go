package models

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (r Route) String() string {
	return fmt.Sprintf("Name: %s & Method: %s & Pattern: %s & HandlerFunc: %s", r.Name, r.Method, r.Pattern, r.HandlerFunc)
}
