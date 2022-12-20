package main

import (
	"fmt"
	"godsvisus"
	"godsvisus/internal/entity"
)

func main() {
	// testList := []int{6, 3, 100, 55, 87, -4}
	// defer godsvisus.ShowArrays([][]int{
	// 	testList,
	// })

	// defer godsvisus.CompareArrays([][]interface{}{
	// 	{6, 3, 100},
	// 	{6, 3, "sekai"},
	// })

	defer godsvisus.CompareLinkedLists([]*entity.Node{
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
