package cap3

import "fmt"

func regressiava(i int64) {
	fmt.Println(i)

	//if i == 0 {
	//	return
	//}

	regressiava(i - 1)
}
