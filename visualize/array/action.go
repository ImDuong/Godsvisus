package array

import (
	"errors"
	"godsvisus/internal/entity"
	"image/color"
	"reflect"
)

func (lay *ArrayLayout) Compare() error {
	if len(lay.components) < 2 {
		return errors.New("there are not enough lists to compare")
	}
	diffNodes := []*entity.ElementWrapper{}
	for i := 0; i < len(lay.components)-1; i++ {
		if len(lay.components[i].Nodes) != len(lay.components[i+1].Nodes) {
			diffNodes = append(diffNodes, lay.components[i].Nodes...)
			diffNodes = append(diffNodes, lay.components[i+1].Nodes...)
			continue
		}
		for j := range lay.components[i].Nodes {
			if !reflect.DeepEqual(lay.components[i].Nodes[j].Data, lay.components[i+1].Nodes[j].Data) {
				diffNodes = append(diffNodes, lay.components[i].Nodes[j], lay.components[i+1].Nodes[j])
			}
		}
	}
	markRed(diffNodes)
	return nil
}

func markRed(nodes []*entity.ElementWrapper) {
	for i := range nodes {
		nodes[i].Shape.FillColor = color.CMYK{0, 94, 100, 0}
	}
}
