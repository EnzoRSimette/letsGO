package main

import "fmt"

// Diferença entre arrays e slices

func arrays() {
	// arrays -> Possuem comprimento fixo e se for feito uma cópia como no exemplo:
	a := [3]int{1, 2, 3}
	b := a    // Cria uma cópia por valor
	b[0] = 99 // NÃO ALTERA A
	fmt.Println(b)

	// slices -> possuem tamanho dinâmico, e se for feito uma cópia como no exemplo:
	// obs: slices são 3 partes, com uma delas sendo um ponteiro para um array
	/* as 3 partes:
	ptr: ponteiro para o primeiro elemento visível
	len: quantos elementos o slice enxerga
	cap: quantos elementos existem no array a partir do ptr
	*/
	c := []int{1, 2, 3, 4, 5} // len=5 cap=5 -> len é o tamanho real e cap é o quando ele aguenta de inf nova
	d := a[1:3]               // len=2 cap=4
	/*
		c.ptr ──>   [ 1 | 2 | 3 | 4 | 5 ]
		d.ptr ────────> [ 2 | 3 | 4 | 5 ]
		                  ↑       ↑
		                d[0]    d[cap-1]
	*/
	//! ELES COMPARTILHAM O MESMO PONTEIRO PARA O ARRAY, ENTÃO:
	d[0] = 99 //! VAI MUDAR c POR QUE O PONTEIRO É O MESMO
	fmt.Println(c[1])

	// Append:
	j := make([]int, 3, 5) // len=3 cap=5
	j = append(j, 4)       // ok
	j = append(j, 5)       // ok
	j = append(j, 6)       // NÃO CABE -> Excedeu o limite do array
	//^ Nesse caso, ele vai criar um novo array com capacidade maior,
	//^ alocando os elementos para o novo array e
	//^ retornar um pointer para o novo array

	// Ex:
	v := []int{1, 2, 3}
	z := v
	v = append(v, 4, 5, 6) // cap estourou — novo array alocado, se separa de z
	z[0] = 99
	fmt.Println(v[0]) // 1 — v e z agora apontam para arrays diferentes

	// nil slice — ptr=nil, len=0, cap=0
	var s []int
	fmt.Println(s == nil) // true
	fmt.Println(len(s))   // 0 — seguro
	s = append(s, 1)      // seguro — append aceita nil slice

	// slice vazio — ptr≠nil, len=0, cap=0
	s = []int{}
	fmt.Println(s == nil) // false

	// make — aloca com len e cap definidos
	s = make([]int, 5)     // len=5, cap=5, todos zero value
	s = make([]int, 0, 10) // len=0, cap=10 — útil quando você sabe o tamanho final

	// Ruim — append realoca várias vezes
	result1 := []int{}
	for i := 0; i < 10000; i++ {
		result1 = append(result1, i)
	}

	// Bom — uma única alocação
	result2 := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		result2 = append(result2, i)
	}
	/*\b
		// Copiar sem copiar o ponteiro:
		a := []int{1, 2, 3, 4, 5}
		b := make([]int, len(a))
		copy(b, a)  // copia os valores, não o ponteiro

		b[0] = 99
		fmt.Println(a)  // [1 2 3 4 5] — não afetado
	\e*/

	/*\b

	Slice de slice em append:
	a := make([]int, 3, 6)  // len=3, cap=6
	b := a[:3]              // compartilha array, cap=6

	b = append(b, 99)       // len=4, cap=6 — escreve em a[3]
	fmt.Println(a[:4])      // [0 0 0 99] — a foi modificado silenciosamente

	Para evitar: use three-index slice a[low:high:max] para limitar o cap do slice derivado:
	b := a[:3:3]  // cap=3 — append vai realocar, não vai tocar em a

	\e*/

}

//* Memory leak por conta de um array grande
func getFirst1(data []byte) []byte {
	return data[:3] // slice de 3 bytes
	// mas o array inteiro (potencialmente GBs) fica na memória
	// porque o slice ainda referencia ele
}

// Correto — copia os dados necessários, libera o array original
func getFirst2(data []byte) []byte {
	result := make([]byte, 3)
	copy(result, data[:3])
	return result
}
