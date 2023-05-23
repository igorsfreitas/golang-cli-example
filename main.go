package main

import (
	"fmt"
	"github/igorsfreitas/golang-cli-example/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Startando Aplicação de Linha de Comando")

	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
