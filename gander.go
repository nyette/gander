package main

type Gander struct {
	Nest string
}

// NewGander returns a new instance of Gander.
func NewGander() Gander {
	g := Gander{Nest: "http://localhost:8000"}
	return g
}
