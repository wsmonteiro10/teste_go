package util

import (
	"regexp"
	"strconv"
)

func ValidaCPF(cpf string) (string, int) {

	if cpf == "NULL" {
		return "00000000000", 1
	}

	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return "", 0
	}

	soma := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (10 - i)
	}
	d1 := 11 - (soma % 11)
	if d1 >= 10 {
		d1 = 0
	}

	soma = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (11 - i)
	}
	d2 := 11 - (soma % 11)
	if d2 >= 10 {
		d2 = 0
	}

	if d1 != int(cpf[9]-'0') || d2 != int(cpf[10]-'0') {
		return "", 0
	}

	// Retorna CPF apenas com números (já está limpo)
	return cpf, 1
}

func ValidaCNPJ(cnpj string) (string, int) {

	if cnpj == "NULL" {
		return "00000000000000", 1
	}

	// Remove tudo que não é número
	re := regexp.MustCompile(`\D`)
	cnpj = re.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return "", 0
	}

	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calcula primeiro dígito verificador
	soma := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * pesos1[i]
	}
	d1 := 11 - (soma % 11)
	if d1 >= 10 {
		d1 = 0
	}

	// Calcula segundo dígito verificador
	soma = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		soma += num * pesos2[i]
	}
	d2 := 11 - (soma % 11)
	if d2 >= 10 {
		d2 = 0
	}

	// Verifica se os dígitos conferem
	if d1 != int(cnpj[12]-'0') || d2 != int(cnpj[13]-'0') {
		return "", 0
	}

	// Retorna apenas os números
	return cnpj, 1
}
