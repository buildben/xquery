package xmlquery

import "strings"

func LoadXML(s string) *Node {
	node, err := Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	return node
}