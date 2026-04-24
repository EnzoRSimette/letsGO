package main

func contain(a []int, v int) bool { // vamos usar int para simplificar
	var result bool
	for _, t := range a {
		if t == v {
			result = true
			continue
		} else {
			result = false
			continue
		}
	}
	return result
}

func unique(nums []int) []int {
	result := make([]int, 0, len(nums)) // Ele aloca com len(nums) por que a quantidade máxima de elementos no resultado sempre
	//  vai ser <= a quantide de elementos em um dos arrays
	for _, v := range nums {
		if contain(result, v) {
			continue
		} else {
			result = append(result, v)
		}
	}
	return result
	// Não sei se você queria um tratamento de erro caso o array venha vazio, mas como não pediu, não fiz
}

func intersect(a, b []int) []int {
	result := make([]int, 0, len(a)) // Pode ser tanto a quanto b
	for _, v := range a {            // aqui pode ser tanto a quanto b
		if contain(b, v) {
			result = append(result, v)
		} else {
			continue
		}
	}
	return result
}
