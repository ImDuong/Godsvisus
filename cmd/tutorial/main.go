package main

import (
	"fmt"

	"github.com/ImDuong/godsvisus"
	"github.com/ImDuong/godsvisus/internal/entity"
)

func main() {
	defer godsvisus.Run()

	testListVis, _ := godsvisus.ShowArrays([][]int{
		{6, 3, 100, 55, 87, -4},
		{6, 3, 100, 55, 87, 3},
	})
	godsvisus.CompareDisplayedArrays(testListVis)

	godsvisus.CompareArrays([][]interface{}{
		{6, 3, 100},
		{6, 3, "sekai"},
	})

	godsvisus.CompareLinkedLists([]*entity.Node{
		{
			Value: 12,
			Next: &entity.Node{
				Value: 3,
				Next: &entity.Node{
					Value: 69,
				},
			},
		},
		{
			Value: 4,
			Next: &entity.Node{
				Value: 3,
			},
		},
	})
	fmt.Println("Hi, because fyne cannot work on another goroutine except main, using defect is the only choice to use this library")
	fmt.Println("What a disposable library :)")
}
