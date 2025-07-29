package etl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"teste_go/util"
)

func Transform(filePath string) {
	// Abre o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\s{2,}`)

	var registros []any
	linha := 0

	for scanner.Scan() {
		linha++

		if linha == 1 {
			continue
		}

		campos := re.Split(strings.TrimSpace(scanner.Text()), -1)

		if len(campos) < 8 {

			for len(campos) < 8 {
				campos = append(campos, "")
			}
		}
		cpf, cpf_valid := util.ValidaCPF(campos[0])
		lf_cnpj, lf_cnpj_valid := util.ValidaCNPJ(campos[7])
		ul_cnpj, ul_cnpj_valid := util.ValidaCNPJ(campos[9])
		record := []any{
			cpf,                   // CPF
			cpf_valid,             // IS_CPF_VALIDO
			campos[2],             // PRIVATE
			campos[3],             // INCOMPLETO
			campos[4],             // ULTIMA_COMPRA
			parseFloat(campos[5]), // TICKET_MEDIO
			parseFloat(campos[6]), // TICKET_ULTIMA_COMPRA
			lf_cnpj,               // LOJA_FREQUENTE
			lf_cnpj_valid,         // IS_LOJA_FREQUENT_VALIDO
			ul_cnpj,               // LOJA_ULTIMA_COMPRA
			ul_cnpj_valid,         // IS_LOJA_ULTIMA_COMPRA_VALIDO
		}
		registros = append(registros, record)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, r := range registros {
		fmt.Printf("%+v\n", r)
	}

}

func parseFloat(s string) float64 {
	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0.0
	}
	return val
}
