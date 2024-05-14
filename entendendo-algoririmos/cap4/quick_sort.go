package cap4

import "math/rand"

func quickSortPivoIndexZero(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	magicNumber := len(arr) + 1
	index := rand.Intn(magicNumber)
	pivo := arr[index]
	menores := []int{}
	maiores := []int{}
	for _, v := range append(arr[:index], arr[index+1:]...) {
		if v <= pivo {
			menores = append(menores, v)
		} else {
			maiores = append(maiores, v)
		}
	}
	return append(append(quickSortPivoIndexZero(menores), pivo), quickSortPivoIndexZero(maiores)...)
}
