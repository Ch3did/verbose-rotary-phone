package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	sites := []string{
		"http://www.google.com",
		"http://www.alura.com",
		"http://www.random-status-code.herokuapp.com/.com",
	}
	//Implementa a seleção de sites
	//	sites = slice(sites)
	for {
		menu()
		comando := leituraInt()
		switch comando {
		case 1:
			monitoramento(sites)
		case 2:
			fmt.Println("Exibindo logs...")
			logs()
		case 0:
			fmt.Println("Até a próxima")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1)
		}

	}

}

func logs() {

}

func monitoramento(sites []string) {
	fmt.Println("Monitorando...")
	fmt.Println("Quantas vezes devo executar?: ")
	vezes := leituraInt()
	for count := 0; count < vezes; count++ {
		for count2, sites := range sites {
			resp, _ := http.Get(sites)
			fmt.Println("Testando site", count2+1, ":"+sites)
			if resp.StatusCode == 200 {
				fmt.Println(sites[count2], "foi carregado com sucesso!")
			} else {
				fmt.Println(sites[count2], "não foi carregados corretamente", resp.StatusCode)
			}
			time.Sleep(time.Second * 2)

		}
	}

	fmt.Println(" ")
}

func leituraInt() int {
	var leitura int
	fmt.Scan(&leitura)
	return leitura

}

func menu() {
	fmt.Println("---------------------------------------------")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do Programa")
}

func slice(sites []string) []string {
	fmt.Print("Me de o site que vc deseja monitorar: ")
	for i := 1; i > 2; i++ {
		s := leituraStrng()
		sites = append(sites, s)
		fmt.Println("Deseja adicionar outro ao monitoramento? (y - Sim | n - Não):")
		choise := leituraStrng()
		switch choise {
		case "n":
			i = 2
		default:
			fmt.Println("Opção Inválida...")
			i = 0
		}

	}
	return sites
}

func leituraStrng() string {
	var leitura string
	fmt.Scan(&leitura)
	return leitura

}
