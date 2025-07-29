package util

import (
	"fmt"
	"regexp"
	"strconv"
)

func ValidaCPF(cpf string) (string, int) {
	// Remove tudo que não é número
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	// Verifica se tem 11 dígitos
	if len(cpf) != 11 {
		return "", 0
	}

	// Calcula o primeiro dígito verificador
	soma := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (10 - i)
	}
	d1 := 11 - (soma % 11)
	if d1 >= 10 {
		d1 = 0
	}

	// Calcula o segundo dígito verificador
	soma = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (11 - i)
	}
	d2 := 11 - (soma % 11)
	if d2 >= 10 {
		d2 = 0
	}

	// Verifica se os dígitos conferem
	if d1 != int(cpf[9]-'0') || d2 != int(cpf[10]-'0') {
		return "", 0
	}

	// Formata o CPF
	cpfFormatado := fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
	return cpfFormatado, 1
}

func ValidaCNPJ(cnpj string) (string, int) {
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

	// Formata: 12.345.678/0001-95
	formatado := fmt.Sprintf("%s.%s.%s/%s-%s",
		cnpj[0:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:14])

	return formatado, 1
}
