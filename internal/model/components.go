package model

import "github.com/a-h/templ"

type Input struct {
	ID              string
	Type            string // defaults to "text"
	Label           string
	Name            string
	Value           string
	Placeholder     string
	Err             string
	Attrs           templ.Attributes
	Class           string
	Icon            templ.Component
	Disabled        bool
	DisabledMessage string
}
