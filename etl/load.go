package etl

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	tableName  = "cliente_compras_raw"
	batchSize  = 1000
	maxWorkers = 4
)

func Salvar_lotes_rapido(db *sql.DB, records [][]any) error {
	chunks := Separar_em_lotes(records, batchSize)

	jobs := make(chan [][]any, len(chunks))
	var wg sync.WaitGroup

	// Lança workers
	for w := 0; w < maxWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range jobs {
				if err := Inserir_em_lote(db, batch); err != nil {
					log.Printf("Erro ao inserir batch: %v", err)
				}
			}
		}()
	}

	for _, batch := range chunks {
		jobs <- batch
	}
	close(jobs)

	wg.Wait()
	return nil
}

func Separar_em_lotes(data [][]any, size int) [][][]any {
	var chunks [][][]any
	for size < len(data) {
		chunks = append(chunks, data[:size])
		data = data[size:]
	}
	if len(data) > 0 {
		chunks = append(chunks, data)
	}
	return chunks
}

func Inserir_em_lote(db *sql.DB, records [][]any) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO cliente_compras_raw (
			CPF, IS_CPF_VALIDO, PRIVATE, INCOMPLETO, ULTIMA_COMPRA,
			TICKET_MEDIO, TICKET_ULTIMA_COMPRA, LOJA_FREQUENTE,
			IS_LOJA_FREQUENT_VALIDO, LOJA_ULTIMA_COMPRA, IS_LOJA_ULTIMA_COMPRA_VALIDO
		) VALUES 
	`

	valStr := ""
	args := []interface{}{}
	argPos := 1

	for _, rec := range records {
		valStr += "("
		for i := 0; i < len(rec); i++ {
			valStr += fmt.Sprintf("$%d", argPos)
			argPos++
			if i < len(rec)-1 {
				valStr += ","
			}
			args = append(args, rec[i])
		}
		valStr += "),"
	}
	valStr = valStr[:len(valStr)-1]

	_, err = tx.Exec(query+valStr, args...)
	if err != nil {
		return err
	}

	return tx.Commit()
}
