package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	JogarDeNovo := true
	var TotasAsTentativas [][]int

	for JogarDeNovo {
		resposta :=
			GerarNumeroAleatorio(1, 100)
		fmt.Println("tente adivinhar o numero entre 1 e 100:")

		var tentativas []int
		palpite := 0
		for palpite != resposta {
			palpite = ObterPalpitesDoUsuario()
			tentativas =
				append(tentativas, palpite)

			if palpite > resposta {
				fmt.Println("o numero é menor!")
			} else if palpite < resposta {
				fmt.Println("o numero é maior!")
			}
		}
		fmt.Printf("parabens, você acertou em %d tentativas!\n", len(tentativas))
		TotasAsTentativas =
			append(TotasAsTentativas, tentativas)

		JogarDeNovo = PediParaJogarDeNovo()
	}

	fmt.Println("numero de tentativas em todas jogadas:")
	printTodasAsTentativas(TotasAsTentativas)

}
func GerarNumeroAleatorio(min, max int) int {

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min

}

func ObterPalpitesDoUsuario() int {
	var palpite int
	fmt.Print("digite seu palpite:")
	fmt.Scan(&palpite)
	return palpite
}

func PediParaJogarDeNovo() bool {
	var JogarDeNovo string
	fmt.Print("deseja jogar novamente? (s\n:")
	fmt.Scan(&JogarDeNovo)
	return JogarDeNovo == "s" ||
		JogarDeNovo == "s"
}

func printTodasAsTentativas(tentativas [][]int) {
	for i, TentativasDoJogo := range tentativas {
		fmt.Printf("jogo %d : %d tentativas\n", i+1, len(TentativasDoJogo))

	}
}
