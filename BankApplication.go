package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"regexp"
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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//fmt.Println("Successfully connected to DB")
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
	// Verificar conexão
	if err := db.Ping(); err != nil {
		fmt.Println("Erro de conexão com o banco de dados:", err)
		return
	}

	psqlscript := `CALL addConta($1, $2, $3, $4)`
	reader := bufio.NewReader(os.Stdin)

	var tipoInt int
	for {
		fmt.Println("\nQue tipo de conta deseja abrir? (Digite o número correspondente)")
		for i := 1; i <= 4; i++ {
			fmt.Printf("%d: %s\n", i, Tipo[i])
		}
		fmt.Print("Escolha: ")

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		tipoInt, err = strconv.Atoi(input)
		if err == nil && tipoInt >= 1 && tipoInt <= 4 {
			break
		}
		fmt.Println("Tipo de conta inválido. Por favor, escolha um número entre 1 e 4.")
	}
	tipoConta := Tipo[tipoInt]

	var nome string
	for {
		fmt.Print("\nQual é o seu nome\nNome: ")
		input, err := reader.ReadString('\n')
		nome = strings.TrimSpace(input)
		if err == nil && nome != "" {
			break
		}
		fmt.Println("O nome não pode ser vazio.")
	}

	var valorEntrada float64
	for {
		fmt.Print("\nInsira um valor de entrada. O valor deve ser entre 50€ e 800€\nValor: ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		valorEntrada, err = strconv.ParseFloat(input, 64)
		if err == nil && valorEntrada >= 50 && valorEntrada <= 800 {
			break
		}
		fmt.Println("Valor inválido. Por favor, insira um valor entre 50€ e 800€.")
	}

	var email string
	for {
		fmt.Print("\nDigite seu e-mail: ")
		input, err := reader.ReadString('\n')
		email = strings.TrimSpace(input)
		if err == nil && isValidEmail(email) {
			break
		}
		fmt.Println("Por favor, insira um e-mail válido.")
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Erro ao iniciar transação:", err)
		return
	}

	_, err = tx.Exec(psqlscript, nome, tipoConta, valorEntrada, email)
	if err != nil {
		tx.Rollback()
		fmt.Println("Erro ao criar conta:", err)
		return
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Erro ao finalizar transação:", err)
		return
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

	fmt.Printf("\nConta criada com sucesso!\nNome: %s\nTipo de Conta: %s\nSaldo Inicial: %.2f€\nEmail: %s\n",
		conta.Nome, conta.TipoConta, conta.Dinheiro, conta.Email)
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(strings.ToLower(email))
}

// TODO: Fix the scanner Bug
func getConta() *Conta {
	var conta *Conta
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o seu email: ")
	scanner.Scan()
	email := scanner.Text()

	isValidEmail(email)

	for _, c := range contas {
		if c.Email == email {
			return conta
		}
		fmt.Println(c)
	}

	fmt.Println("Nenhuma conta encontrada com este email")
	return nil
}

func main() {
	db := connectDB()
	defer db.Close()
	fmt.Println("Bem-vindo ao sistema bancário!")
	fmt.Println("Escolha uma opção:")
	fmt.Println("1: Abrir nova conta")
	fmt.Println("2: Ver Conta")
	fmt.Println("0: Sair")

	var escolha int
	fmt.Scan(&escolha)

	switch escolha {
	case 1:
		abrirConta(db)
	case 2:
		getConta()
	case 0:
		fmt.Println("Saindo do sistema.")
	default:
		fmt.Println("Opção inválida.")
	}
}

func timeSleep() {
	time.Sleep(2 * time.Second)
}
