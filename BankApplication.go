package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Tipo = map[int]string{
	1: "Normal",
	2: "Familia",
	3: "Empresa",
	4: "Estudante",
}

var contas []Conta

type Conta struct {
	ID             int
	Nome           string
	TipoConta      string
	ValordeEntrada float64
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
		return
	}
	if NovoDinheiro > 800 || NovoDinheiro < 0 {
		fmt.Println("O Valor está errado")
		return
	}
	if NovoDinheiro < 20 {
		fmt.Println("O Valor Minimno de Deposito é de 20")
		return
	}
	conta.Dinheiro += NovoDinheiro
	fmt.Printf("Valor atualizado da conta: %.2f€\n", conta.Dinheiro)
}

func removeDinheiro(conta *Conta) float64 {
	var dinheiro float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Qual é o Valor da Conta que deseja Tirar")
	scanner.Scan()
	dinheiro, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Valor Errado")
		return 0
	}
	if conta.Dinheiro < dinheiro || dinheiro < 0 {
		fmt.Println("Esse valor não é autorizado")
		return 0
	}
	conta.Dinheiro -= dinheiro
	fmt.Printf("Valor atualizado da conta: %.2f€\n", conta.Dinheiro)

	return dinheiro
}

func abrirConta(Nome string, tipoD string, valorE float64) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Que tipo de conta deseja abrir? (Digite o número correspondente)")
	fmt.Println("1: Normal")
	fmt.Println("2: Familia")
	fmt.Println("3: Empresa")
	fmt.Println("4: Estudante")
	scanner.Scan()
	tipoInput := scanner.Text()

	tipoInt, err := strconv.Atoi(tipoInput)
	if err != nil || tipoInt < 1 || tipoInt > 4 {
		fmt.Println("Tipo de conta inválido. Operação cancelada.")
		return
	}

	fmt.Print("Insira um valor de entrada. O valor deve ser no mínimo de 50€ \nValor: ")
	scanner.Scan()
	valorEntradaInput := scanner.Text()

	valorEntrada, err := strconv.ParseFloat(valorEntradaInput, 64)
	if err != nil || valorEntrada < 50 {
		fmt.Println("Valor de entrada inválido. Operação cancelada.")
		return
	}

	id := len(contas) + 1

	conta := Conta{
		ID:             id,
		Nome:           Nome,
		TipoConta:      Tipo[tipoInt],
		ValordeEntrada: valorE,
		Dinheiro:       valorE,
	}

	contas = append(contas, conta)

	fmt.Printf("\nConta criada com sucesso! Tipo de conta: %s, Tipo de Conta: %s\n", conta.Nome, conta.TipoConta)

}

func main() {

}
