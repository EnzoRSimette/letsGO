package main

import "fmt"

// Em go, assim como em js, funções são first-class
// na prática, funcao1 é um ponteiro para uma função anonima
var funcao1 = func() { // pode usar := dentro de outra funcao
	fmt.Println("Função atribuida a var")
}

func makeCounter() func() int /* aqui, estamos retornando uma outra função */ {
	count := 0          // esta variável vai para o heap
	return func() int { // esta função captura count pois está dentro da função que contém count
		count++
		return count
		/*
			O que acontece na memória — isso é importante:
			! count normalmente viveria na stack de makeCounter.
			? Mas como a função anônima retornada ainda referencia count após makeCounter retornar,
			* o compilador detecta isso via escape analysis e move count para o heap.
			~ Isso é uma alocação. Closures que capturam variáveis têm custo.
			Em hot paths (código executado milhões de vezes), isso importa.
			Você vai medir exatamente isso com go build -gcflags='-m' na Fase 3.
		*/
	}
}

func funcoes() {
	counter := makeCounter()
	counter() // 1
	counter() // 2
	counter() // 3
}

// Naked returns!
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // "naked return" — vai retornar o que é definido em () -> O nome é autoexplícito
}

type Order struct { // dead code apenas para não crashar tudo
}

// BOM: função curta, nomes adicionam documentação real
func multiply(a, b float64) (result float64, err error) { /*...*/ return }

// RUIM: função longa com naked returns — o leitor perde o rastreio
func processOrder(o Order) (id int, total float64, tax float64, err error) {
	// 50 linhas...
	return // o que está sendo retornado aqui?
}

//^ Em go, funções podem retornar mais de um valor! (isso faz com que não precise existir try catch)
func divide(a, b float64) (float64, error) {
	/*
		Por baixo dos panos:
		Em arquiteturas x86-64, funções retornam valores via registradores — RAX, RDX para os dois primeiros valores.
		O compilador Go aloca os retornos nesses registradores diretamente.
		Múltiplos retornos simples não custam alocação de heap — são registradores ou posições na stack, dependendo dos tipos.
		Isso contrasta com linguagens que retornam tuplas como objetos heap-alocados.
		Em Go, (float64, error) são dois valores na stack.
		Zero overhead de alocação para o caso comum.

		ex: |float64| |error| -> funcao -> |retornofloat| |retornoerro| (Mesmos registros!)
	*/
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero!")
	}
	return a / b, nil // onde nil é zero-value pro tipo
}

// T: Filosofia do (value, error) -> em JS e Java, erros são excecões, o fluxo salta para um catch
// JS: Você não sabe, só lendo a assinatura, que isso pode falhar
// function parseUser(data) { ... }
// Go: Você sabe imediatamente que isso pode falhar — está na assinatura
// func parseUser(data []byte) (User, error) {...}
//t -> O erro é um valor comum, não um mecanismo de controle de fluxo.
//t -> Consequência direta: o compilador te força a lidar com ele ou descartá-lo conscientemente.
/*
	user, err := parseUser(data)
	if err != nil {
		return fmt.Errorf("parseUser: %w", err)  // propaga com contexto
	}
	-> aqui, user é garantidamente válido
*/
