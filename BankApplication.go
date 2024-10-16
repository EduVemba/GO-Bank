package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tipo = map[int]string

type Conta struct {
	Tipo           Tipo
	ValordeEntrada int64
	Dinheiro       float64
}

func addDinheiro(conta *Conta) {
	var NovoDinheiro float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Qual é o Valor da Conta que deseja Meter")
	scanner.Scan()
	NovoDinheiro, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Valor Errado")
	}
	if NovoDinheiro > 800 && NovoDinheiro < 0 {
		fmt.Println("O Valor está errado")
	}
	conta.Dinheiro = NovoDinheiro
	fmt.Printf("Valor atualizado da conta: %.2f€\n", conta.Dinheiro)
}

func removeDinheiro(conta *Conta) float64 {
	var dinheiro float64

}

func abrirConta(conta Conta) Conta {
	if conta.ValordeEntrada >= 49 {
		fmt.Println("Valor De entrada tem de ser no minimo 50€")
	}
	NovaConta := conta

	return NovaConta
}

func main() {

}
