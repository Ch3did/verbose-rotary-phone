package utils

import "fmt"

func LeituraInt() int {
	var leitura int
	fmt.Scan(&leitura)
	return leitura

}

func LeituraStrng() string {
	var leitura string
	fmt.Scan(&leitura)
	return leitura

}

func Menu() {
	fmt.Println("---------------------------------------------")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do Programa")
}
