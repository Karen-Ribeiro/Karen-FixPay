package main

import (
	"fmt"
	"os"

)

func main(){

	for {

		exibirMenu()
		opcao := lerComando()
		switch opcao {
		case 1:

			for i := 0; i != -1; i++{
				fmt.Println("----- Inserindo um novo contato -----")

				registraNomes(i)
				registraNumeros(i)

				fmt.Println("")
				fmt.Println("Você deseja inserir um novo contato?")
				fmt.Println("Digite S para confirmar ou N para negar.")

				var letraDeControle string
				fmt.Scan(&letraDeControle)

				if letraDeControle == "S"{
					i = i
				} else
				{
					i = -1
				}
			}

		case 2:
			fmt.Println("----- Exibindo Contatos -----")

		case 0:
			fmt.Println("----- Saindo do programa -----")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida!\n Fechando o programa.")
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

func registraNomes(i int){
	var nomes [100]string
	var nome string
	fmt.Println("Digite o nome do contato:")
	fmt.Scan(&nome)

	nomes[i] = nome

	arquivo, err := os.OpenFile("nomes.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println(err)
	}

	arquivo.WriteString("Nome:"  + nome + " - ")
	arquivo.Close()
}

func registraNumeros(i int){
	var numeros [100]string
	var numero string

	fmt.Println("Digite o numero do contato:")
	fmt.Scan(&numero)

	numeros[i] = numero

	arquivo, err := os.OpenFile("nomes.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println(err)
	}

	arquivo.WriteString("Numero:"  + numero + "\n")
	arquivo.Close()
}
