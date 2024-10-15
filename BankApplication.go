package main

import "fmt"

type Tipo = map[int]string

type Conta struct {
	Tipo           Tipo
	ValordeEntrada int64
}

func abrirConta(conta Conta) Conta {
	if conta.ValordeEntrada >= 49 {
		fmt.Println("Valor De entrada tem de ser no minimo 50€")
	}
	NovaConta := conta

	return NovaConta
}

func AbrirConta() {}

func main() {

}
