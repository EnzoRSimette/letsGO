package nao_funcionais

import "fmt"

func variaveis() {
	// Em go, uma var é um nome para
	// uma região específica na memória,
	// scom tamanho fixo e conhecido na comp
	// obs: quando você declara sem atribuir valor ele fica como zero-value
	// ex:
	//  |_| |_| |_| |_| (começo)
	// -> n1 := "a" (OBS: := só pode ser usado em funções)
	// 	|a| |_| |_| |_|

	// var b int -> aloca 8 bits de memória em alguma "caixinha", decidido pelo compilador
	//! zero-values: Valores padrão para quando não tem definição da variável

	x := 1
	x := 2 // ERRO: no new variables on left side of :=

	y := 1
	y, z := 2, 3

	/*
		+-----------+------------+------------------+
		| Tipo      | Zero Value | Tamanho (64-bit) |			OBS: x := 1
		+-----------+------------+------------------+				 x := 2 x -> ERRO: no new variables on left side of :=
		| int,int64 | 0          | 8 bytes          |				 ^ Ocorre pois estamos redefinindo a var
		| int32     | 0          | 4 bytes          |
		| float64   | 0.0        | 8 bytes          |				 x := 1
		| bool      | false      | 1 byte           |				 x, y := 2, 3 -> // OK: y é nova, x é reatribuída
		| string    | ""         | 16 bytes*        |
		| pointer   | nil        | 8 bytes          |			VAR BLOCK:
		| slice     | nil        | 24 bytes*        |				var (
		| map       | nil        | 8 bytes          |					host    string = "localhost"
		| struct    | cada campo | soma dos campos  |					port    int = 8000
		|           | no zero    |                  |					timeout int = 30
		|           | value      |                  |				)
		+-----------+------------+------------------+
	*/

	/*
		! $$$$$$$$$$$$$$$$$$$$
		! $ MUITO IMPORTANTE $
		! $$$$$$$$$$$$$$$$$$$$

			int8    // -128 a 127            — 1 byte
			int16   // -32768 a 32767        — 2 bytes
			int32   // -2B a 2B              — 4 bytes
			int64   // -9.2Q a 9.2Q          — 8 bytes
			int     // depende da plataforma — 4 ou 8 bytes

			uint8   // 0 a 255              — 1 byte  (byte é alias de uint8)
			uint16  // 0 a 65535            — 2 bytes
			uint32  // 0 a 4B               — 4 bytes
			uint64  // 0 a 18.4Q            — 8 bytes
			uint    // depende da plataforma

			Por que isso importa no hardware:
			A CPU tem registradores de 64 bits em arquiteturas modernas (x86-64, ARM64).
			Operações em int64 são nativas — custam exatamente uma instrução de CPU.
			Operações em int8 dentro de loops podem custar mais
			porque o processador precisa aplicar máscaras para truncar o resultado aos
			8 bits corretos.
	*/

	// Use int por padrão para índices e contadores
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println(i)
	}

	// Use int64 para IDs de banco de dados, timestamps Unix
	var userID int64

	// Use tipos menores apenas quando o tamanho da estrutura importa
	// (structs com milhões de instâncias, serialização binária)
	type Pixel struct {
		R, G, B uint8 // 3 bytes vs 24 bytes com int
	}

	// constantes -> Elas só existem na compilação, por isso, não ocupam memória nem ram depois da compilação
	// ^ Muito bom para performance
	// Tipadas vs Não-Tipadas
	const x = 42      // Assume o tipo conforme contexto
	const y int = 42  // bem definido
	var a float64 = x // Pode, já que x não é tipado
	var b float64 = y // Não pode, já que y é int
}
