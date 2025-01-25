package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Data     interface{}
	Children []TreeNode
}

func (t *TreeNode) String(level int) string {
	indent := strings.Repeat(" ", level)
	result := fmt.Sprintf("%s%s\n", indent, t.Data)
	for _, child := range t.Children {
		result += child.String(level + 1)
	}
	return result
}

func (t *TreeNode) addChild(child TreeNode) {
	t.Children = append(t.Children, child)
}

func main() {
	tree := TreeNode{Data: "Drinks"}
	cold := TreeNode{Data: "Cold"}
	hot := TreeNode{Data: "Hot"}
	tree.addChild(cold)
	tree.addChild(hot)
	fmt.Println(tree.String(0))
}
