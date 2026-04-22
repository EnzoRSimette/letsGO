package nao_funcionais

import "strings"

func funcionamentoString(words []string) {
	// Strings não são ponteiros simples, é uma estrutura de campos
	/*
		string header (16 bytes):
		┌─────────────────┬──────────┐
		│  ptr (*byte)    │  len     │
		│  8 bytes        │  8 bytes │
		└─────────────────┴──────────┘
		     │
		     ▼
		[ H | e | l | l | o ]  ← dados imutáveis na memória


		^ CONSEQUÊNCIA:
		todo: TODA string é imutável, isso faz
		todo: Com que toda string tenha o mesmo
		todo: tamanho na memória, quando você
		todo: concatena usando +, você cria uma nova string
	*/

	var s string = "Hello"

	// Como usar a consequência ao seu favor:

	// errado
	result1 := ""
	for _, word := range words {
		result1 += word // Nova alocação a cada interação
	}

	// correto
	var b strings.Builder // Isso aqui aloca espaço específico para concatenação
	for _, word := range words {
		b.WriteString(word) // Vai concatenar as strings em b
	}
	result2 := b.String() // Transforma os bytes de b em string
}
