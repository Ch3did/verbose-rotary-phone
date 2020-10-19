package handlers

import (
	"fmt"
	"net/http"
	"time"

	"main.go/utils"
)

func Logs() {

}

func Monitoramento(sites []string) {
	fmt.Println("Monitorando...")
	fmt.Println("Quantas vezes devo executar?: ")
	vezes := utils.LeituraInt()
	for count := 0; count < vezes; count++ {
		for count2 := 0; count2 < len(sites); count2++ {
			resp, _ := http.Get(sites[count2])
			//algo errado com essa parte do código
			fmt.Println("Testando site", count2+1, ":"+sites[count2])
			if resp.StatusCode == 200 {
				fmt.Println(sites[count2], "foi carregado com sucesso!")
			} else {
				fmt.Println(sites[count2], "não foi carregados corretamente", resp.StatusCode)
			}
			time.Sleep(time.Second * 1)
		}

	}

	fmt.Println(" ")
}

func Slice(sites []string) []string {
	var boole bool = true
	var sf string
	fmt.Print("Me de o site que vc deseja monitorar: ")
	s := utils.LeituraStrng()
	sf = "https://www." + s + ".com.br"
	sites = append(sites, sf)
	for boole == true {
		fmt.Print("Deseja adicionar outro ao monitoramento? (y - Sim | n - Não):")
		fmt.Print(" : ")
		choise := utils.LeituraStrng()
		switch choise {
		case "n":
			fmt.Println(sites)
			boole = false
		case "y":
			fmt.Print("Me dê o próximo: ")
			s = utils.LeituraStrng()
			sf = "https://www." + s + ".com.br"
			sites = append(sites, sf)
		default:
			fmt.Println("Opção Inválida...")
		}

	}
	return sites
}
