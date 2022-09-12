package main

import (
	"bufio"
	"fmt"
	"os"

)

func main(){


	for {

		exibirMenu()
		opcao := lerComando()
		switch opcao {
		case 1:
			i := 0
			for i != -1{
				fmt.Println("\n----- Inserindo um novo contato -----")

				registraNomes()
				registraNumeros()

				fmt.Println("")
				fmt.Println("Você deseja inserir um novo contato?")
				fmt.Println("Digite S para confirmar ou N para negar.")

				var letraDeControle string
				fmt.Scan(&letraDeControle)

				if letraDeControle == "N" || letraDeControle == "n"{
					i = -1
					fmt.Println("Saindo do loop")
				} else {
					fmt.Println("Continuando")
				}
			}

		case 2:
			fmt.Println("\n----- Exibindo Contatos -----")
			lerContatos()

		case 0:
			fmt.Println("\n----- Saindo do programa -----")
			os.Exit(0)

		default:
			fmt.Println("\nOpção inválida!\n Fechando o programa.")
			os.Exit(-1)

		}
	}
}

func exibirMenu(){

	fmt.Println(" ------ Agenda ------")
	fmt.Println("1 - Cadastrar um novo contato")
	fmt.Println("2 - Exibir os números existentes")
	fmt.Println("0 - Sair da agenda")
}

func lerComando() int{
	var opcao int
	fmt.Println("Digite um número: ")
	fmt.Scan(&opcao)
	return opcao
}

func registraNomes(){
	var nome string
	fmt.Println("Digite o nome do contato:")
	fmt.Scan(&nome)


	arquivo, err := os.OpenFile("contatos.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println(err)
	}

	arquivo.WriteString("Nome: "  + nome + " - ")
	arquivo.Close()
}

func registraNumeros(){
	var numero string

	fmt.Println("Digite o numero do contato:")
	fmt.Scan(&numero)

	arquivo, err := os.OpenFile("contatos.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println(err)
	}

	arquivo.WriteString("Numero: "  + numero + "\n")
	arquivo.Close()
}
//Funcao retirada da internet
//https://www.delftstack.com/pt/howto/go/how-to-read-a-file-line-by-line-in-go/
func lerContatos() {

	//Abre o arquivo
	arquivo, err := os.Open("contatos.txt")

	if err != nil{
		fmt.Println("Erro:", err)
	}

	fileScanner := bufio.NewScanner(arquivo)

	//Lendo linha por linha
	for fileScanner.Scan(){

		fmt.Println(fileScanner.Text())
	}
	arquivo.Close()
}
