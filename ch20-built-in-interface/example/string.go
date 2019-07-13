package example

import "fmt"

type Course struct {
	Title    string
	SubTitle string
}

func (c *Course) String() string {
	return fmt.Sprintf("[Course]{Title=%s,SubTitle=%s}", c.Title, c.SubTitle)
}
