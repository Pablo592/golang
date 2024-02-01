package arreglos_slices

import "fmt"

var tabla_slice []int = []int{5, 6, 7}
var arreglo [10]int = [10]int{5, 6, 7}

func MuestroSlice() {
	fmt.Println(tabla_slice)

	porcion := arreglo[3:]   //slice de 3 al final
	porcion2 := arreglo[:5]  //slice de 0 a 5
	porcion3 := arreglo[6:7] //slice de 6 a 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20) // Va a tener 5 elementos y una capacidad de 20

	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0) // Va a tener 0 elementos y una capacidad de 0

	//for i := 0; i < 100; i++ {
	//	nums = append(nums, i)
	//}

	nums = append(nums, 10, 55, 44)
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
}
