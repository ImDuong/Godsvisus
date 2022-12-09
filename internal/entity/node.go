package entity

import "fyne.io/fyne/v2/canvas"

type (
	Node struct {
		Value interface{}
		Next  *Node
	}

	NodeWrapper struct {
		Data      *Node
		Component *canvas.Circle
		Text      *canvas.Text
		Next      *NodeWrapper
		Prev      *NodeWrapper
	}

	LinkedList struct {
		Root        *NodeWrapper
		Connections []*canvas.Line
	}
)

// TODO: add unit test
func NewNodeWrapper(node *Node) *NodeWrapper {
	rootNodeWrapper := &NodeWrapper{}
	if node == nil {
		rootNodeWrapper = nil
	} else {
		curNode := node
		curNodeWrapper := rootNodeWrapper
		curNodeWrapper.Data = curNode
		curNode = curNode.Next
		for curNode != nil {
			// init a new node
			curNodeWrapper.Next = &NodeWrapper{
				Data: curNode,
			}
			// attach new node to old node
			curNodeWrapper.Next.Prev = curNodeWrapper

			// traverse to new node
			curNodeWrapper = curNodeWrapper.Next
			curNode = curNode.Next
		}
	}
	return rootNodeWrapper
}
