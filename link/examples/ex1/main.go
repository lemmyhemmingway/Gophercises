package main

import (
	"fmt"
	"strings"

	"../../../link"
)

var templateHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to other page</a>
  <a href="/another-page">A link to another page</a>
</body>
</html>
`

func main() {
	s := strings.NewReader(templateHtml)
	links, err := link.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(links)
}
