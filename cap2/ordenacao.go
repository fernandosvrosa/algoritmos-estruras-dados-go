package cap2

func buscaMenor(arr []int) (int, int) {
	menor := arr[0]
	menorIndex := 0
	for i, v := range arr {
		if v < menor {
			menor = v
			menorIndex = i
		}
	}
	return menor, menorIndex
}

func ordenacaoPorSelecao(arr []int) []int {
	novoArr := []int{}
	count := len(arr)
	for i := 0; i < count; i++ {
		menor, menorIndex := buscaMenor(arr)
		novoArr = append(novoArr, menor)
		arr = append(arr[:menorIndex], arr[menorIndex+1:]...)
	}
	return novoArr
}
