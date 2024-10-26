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
	"strings"
	"time"
)

func connectDB() *sql.DB {
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

	return db

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
	Email     string
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
func abrirConta(db *sql.DB) {
	psqlscript := `CALL addConta($1, $2, $3, $4)`
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Que tipo de conta deseja abrir? (Digite o número correspondente)")
	fmt.Println("1: Normal")
	fmt.Println("2: Familia")
	fmt.Println("3: Empresa")
	fmt.Println("4: Estudante")
	fmt.Print("Escolha: ")
	//TODO: FIX: por algum motivo esta a ler um valor vazio
	tipoInput, _ := reader.ReadString('\n')
	tipoInput = strings.TrimSpace(tipoInput)

	tipoInt, err := strconv.Atoi(tipoInput)
	if err != nil || tipoInt < 1 || tipoInt > 4 || Tipo[tipoInt] == "" {
		fmt.Println("Tipo de conta inválido. Operação cancelada.")
		return
	}
	tipoConta := Tipo[tipoInt]

	fmt.Print("Qual é o seu nome \nNome: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)
	if nome == "" {
		fmt.Println("O nome não pode ser vazio.")
		return
	}

	fmt.Print("Insira um valor de entrada. O valor deve ser no mínimo de 50€ \nValor: ")
	valorEntradaInput, _ := reader.ReadString('\n')
	valorEntradaInput = strings.TrimSpace(valorEntradaInput)

	valorEntrada, err := strconv.ParseFloat(valorEntradaInput, 64)
	if err != nil || valorEntrada < 50 || valorEntrada > 800 {
		fmt.Println("Valor de entrada inválido. Operação cancelada.")
		return
	}

	fmt.Print("Digite seu e-mail: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		fmt.Println("O e-mail não pode ser vazio.")
		return
	}

	_, err = db.Exec(psqlscript, nome, tipoConta, valorEntrada, email)
	if err != nil {
		log.Fatal("Erro ao executar a procedure:", err)
	}

	var id int
	for _, conta := range contas {
		if conta.ID >= id {
			id = conta.ID + 1
		}
	}

	conta := Conta{
		ID:        id,
		Nome:      nome,
		TipoConta: tipoConta,
		Dinheiro:  valorEntrada,
		Email:     email,
	}
	contas = append(contas, conta)

	fmt.Printf("\nConta criada com sucesso! Nome: %s, Tipo de Conta: %s\n", conta.Nome, conta.TipoConta)
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
	db := connectDB()
	defer db.Close()
	fmt.Println("Bem-vindo ao sistema bancário!")
	fmt.Println("Escolha uma opção:")
	fmt.Println("1: Abrir nova conta")
	fmt.Println("0: Sair")

	var escolha int
	fmt.Scan(&escolha)

	switch escolha {
	case 1:
		abrirConta(db)
	case 0:
		fmt.Println("Saindo do sistema.")
	default:
		fmt.Println("Opção inválida.")
	}
}

func timeSleep() {
	time.Sleep(2 * time.Second)
}
