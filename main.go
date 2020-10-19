package main

import (
	"fmt"
	"os"

	"main.go/handlers"
	"main.go/utils"
)

func main() {
	var sites []string
	//Implementa a seleção de sites
	sites = handlers.Slice(sites)
	for {
		utils.Menu()
		comando := utils.LeituraInt()
		switch comando {
		case 1:
			handlers.Monitoramento(sites)
		case 2:
			fmt.Println("Exibindo logs...")
			handlers.Logs()
		case 0:
			fmt.Println("Até a próxima")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido")
			os.Exit(-1)
		}

	}
}
