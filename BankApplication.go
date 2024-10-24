package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

func connectDB() {
	err := godotenv.Load("bd_connect.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB")

}

var Tipo = map[int]string{
	1: "Normal",
	2: "Familia",
	3: "Empresa",
	4: "Estudante",
}

var contas []Conta

type Conta struct {
	ID        int
	Nome      string
	TipoConta string
	Dinheiro  float64
}

func addDinheiro() {
	var conta *Conta
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

func removeDinheiro() float64 {
	var conta *Conta
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

func abrirConta() {

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

	fmt.Print("Qual é o seu nome \nNome:")
	scanner.Scan()
	ValorNome := scanner.Text()

	nome := ValorNome
	if nome == "" {
		fmt.Println("O Nome não poder ser vazio")
		return
	}

	fmt.Print("Insira um valor de entrada. O valor deve ser no mínimo de 50€ \nValor: ")
	scanner.Scan()
	valorEntradaInput := scanner.Text()

	valorEntrada, err := strconv.ParseFloat(valorEntradaInput, 64)
	if err != nil || valorEntrada < 50 || valorEntrada > 800 {
		fmt.Println("Valor de entrada inválido. Operação cancelada.")
		return
	}

	var id int
	for _, conta := range contas {
		id = conta.ID + 1
	}
	id++

	conta := Conta{
		ID:        id,
		Nome:      nome,
		TipoConta: Tipo[tipoInt],
		Dinheiro:  valorEntrada,
	}

	contas = append(contas, conta)

	fmt.Printf("\nConta criada com sucesso! Tipo de conta: %s, Tipo de Conta: %s\n", conta.Nome, conta.TipoConta)
}

func getConta() *Conta {
	var id int
	var nome string
	var conta *Conta
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Qual a o ID ou o Nome da conta que pretende encontrar")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		nome = input
	}
	arr := contas
	low := 0
	high := len(arr) - 1
	if err == nil {
		for low <= high {
			mid := low + (high-low)/2
			if arr[mid].ID == id {
				return conta
			}
			if arr[mid].ID < id {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return nil
	}
	for _, c := range contas {
		if c.Nome == nome {
			conta = &c
			return conta
		}
	}
	fmt.Println("Nenhuma Conta encontrada")
	return nil
}

func main() {
	connectDB()
	/*
			scanner := bufio.NewScanner(os.Stdin)

			for {
				// Menu principal
				fmt.Println("\nBem Vindo Ao Vemba´s Bank")
				fmt.Println(`
		     1. Abrir conta
		     2. Ver Conta
		     3. Remover Dinheiro
		     4. Adicionar Dinheiro
		     5. Sair
		        `)

				// Solicita a opção do usuário
				fmt.Print("Escolha uma opção: ")
				scanner.Scan()

				opcao, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Opção inválida. Por favor, tente novamente.")
					continue
				}

				switch opcao {
				case 1:
					abrirConta()
				case 2:
					getConta()
				case 3:
					removeDinheiro()
				case 4:
					addDinheiro()
				case 5:
					fmt.Println("Obrigado por usar o Vemba's Bank. Adeus!")
					return
				default:
					fmt.Println("Opção inválida. Por favor, escolha uma opção entre 1 e 5.")
				}
			}

	*/

}

func timeSleep() {
	time.Sleep(2 * time.Second)
}
