package linkedlist

import (
	"errors"
	"image/color"
	"reflect"

	"github.com/ImDuong/godsvisus/internal/entity"
)

func (lay *LinkedListLayout) Compare() error {
	if len(lay.components) < 2 {
		return errors.New("there are not enough lists to compare")
	}
	diffNodes := []*entity.NodeWrapper{}
	for i := 0; i < len(lay.components)-1; i++ {
		curNode0 := lay.components[i].Root
		curNode1 := lay.components[i+1].Root
		for curNode0 != nil || curNode1 != nil {
			if curNode0 == nil {
				diffNodes = append(diffNodes, curNode1)
				break
			}
			if curNode1 == nil {
				diffNodes = append(diffNodes, curNode0)
				break
			}
			if !reflect.DeepEqual(curNode0.Data.Value, curNode1.Data.Value) {
				diffNodes = append(diffNodes, curNode0, curNode1)
			}
			curNode0 = curNode0.Next
			curNode1 = curNode1.Next
		}
	}
	markRed(diffNodes)
	return nil
}

func markRed(nodes []*entity.NodeWrapper) {
	for i := range nodes {
		nodes[i].Shape.FillColor = color.CMYK{0, 94, 100, 0}
	}
}
