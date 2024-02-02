package main

import "fmt"

func main() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	fmt.Println("len:", len(order), cap(order))
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))

	sli4 := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	sli := sli4[:5:5]
	fmt.Println(sli)
	sli2 := sli4[5:]
	fmt.Println("sli2: ", len(sli2), sli2)
	sli3 := sli2[:5:5]
	fmt.Println("sli3: ", len(sli3))

}
