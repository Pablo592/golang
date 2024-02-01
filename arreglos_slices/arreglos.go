package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{5, 6, 7}

var matriz [5][7]int = [5][7]int{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}, {22, 23, 24, 25, 26, 27, 28}, {29, 30, 31, 32, 33, 34, 35}}

func MuestroArreglos() {

	tabla[0] = 1
	tabla[5] = 15

	tabla2 := [10]string{"Mexico", "Colombia", "Argentina", "Peru"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(matriz[i][j], "\t")
		}
		fmt.Println("")
	}

}
