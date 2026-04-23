/*
    {
        "key": "backspace",
        "command": "-cosmicCursor.deleteLeft",
        "when": "editorTextFocus && !editorReadonly",
    },
    {
        "key": "backspace",
        "command": "-deleteLeft",
        "when": "textInputFocus",
    },

	config reserva ^
*/

package main

import (
	"fmt"
	"math"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return (a + b), nil
	case "-":
		return (a - b), nil
	case "*":
		return (a * b), nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by 0")
		}
		return (a / b), nil
	case "**":
		return (math.Pow(a, b)), nil
	case "%":
		return (math.Mod(a, b)), nil
	default:
		return 0, fmt.Errorf("invalid operation")
	}
}

func makeCalculator() (func(float64, float64, string) (float64, error), func() []string) { // Na prática, isso aqui é uma classe
	var h []string
	calc := func(a, b float64, op string) (float64, error) { // Função para calcular e retornar o resultado
		result, err := calculate(a, b, op) // Calcula com base nos parâmetros
		if err != nil {                    // Verificação de erros
			return 0, err
		}
		entry := fmt.Sprintf("%s(%.2f, %.2f) = %.2f", op, a, b, result) // Salvamos o resultado como entry para poder colocar em h
		h = append(h, entry)
		return result, nil
	}
	getHistory := func() []string {
		return h // Retorna o array cheio
	}
	return calc, getHistory // Retorna as funções criadas para poderem ser acessadas
}
