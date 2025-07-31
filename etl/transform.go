package etl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"teste_go/util"
	"time"
)

func Transform(filePath string, tipo string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("erro ao abrir o arquivo %s: %v", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var todosRegistros [][]any
	linha := 0

	for scanner.Scan() {
		linha++
		if linha == 1 {
			continue // pula cabeçalho, se houver
		}

		// Divide por qualquer quantidade de espaços (robusto)
		campos := strings.Fields(scanner.Text())
		//fmt.Println(campos)

		// Garante 8 campos no total
		for len(campos) < 8 {
			campos = append(campos, "")
		}

		cpf, cpf_valid := util.ValidaCPF(campos[0])

		// Verifica se há CNPJ válido nas posições esperadas (6 e 7)
		lf_cnpj, lf_cnpj_valid := util.ValidaCNPJ(campos[6])
		ul_cnpj, ul_cnpj_valid := util.ValidaCNPJ(campos[7])

		// Data pode estar no campo 3
		ultima_compra := strings.TrimSpace(campos[3])
		if ultima_compra == "" || ultima_compra == "NULL" {
			ultima_compra = time.Now().Format("2006-01-02")
		}

		registro := []any{
			cpf,                         // CPF
			cpf_valid,                   // IS_CPF_VALIDO
			parseNullableInt(campos[1]), // PRIVATE
			parseNullableInt(campos[2]), // INCOMPLETO
			ultima_compra,               // ULTIMA_COMPRA
			parseFloat(campos[4]),       // TICKET_MEDIO
			parseFloat(campos[5]),       // TICKET_ULTIMA_COMPRA
			lf_cnpj,                     // LOJA_FREQUENTE
			lf_cnpj_valid,               // IS_LOJA_FREQUENTE_VALIDO
			ul_cnpj,                     // LOJA_ULTIMA_COMPRA
			ul_cnpj_valid,               // IS_LOJA_ULTIMA_COMPRA_VALIDO
		}

		todosRegistros = append(todosRegistros, registro)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if tipo == "rapido" {
		fmt.Println("Iniciando ETL rápido...")
		if err := SalvarLotes(todosRegistros); err != nil {
			log.Fatalf("Erro ao salvar registros: %v", err)
		}
	}

	if tipo == "ultrarapido" {
		fmt.Println("Iniciando ETL ultrarrápido...")
		if err := SalvarLotesfast(todosRegistros); err != nil {
			log.Fatalf("Erro ao salvar registros: %v", err)
		}
	}

	fmt.Printf("Total de registros processados: %d\n", len(todosRegistros))
}

func parseFloat(s string) float64 {
	s = strings.Replace(s, ",", ".", -1)
	if strings.ToUpper(strings.TrimSpace(s)) == "NULL" || s == "" {
		return 0
	}
	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0
	}
	return val
}

func parseNullableInt(s string) interface{} {

	s = strings.Replace(s, ",", ".", -1)

	if strings.ToUpper(strings.TrimSpace(s)) == "NULL" || s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
