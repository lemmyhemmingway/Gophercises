package main

import (
	"fmt"
	"strings"

	"../../../link"
)

var templateHtml = `
<html>
<body>
<a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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
