package etl

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db           *sql.DB
	POSTGRES_URL string
)

const (
	batchSize  = 1000
	maxWorkers = 4
)

func init() {
	POSTGRES_URL = os.Getenv("POSTGRES_URL")
	if POSTGRES_URL == "" {
		POSTGRES_URL = "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
	}

	var err error
	db, err = sql.Open("postgres", POSTGRES_URL)
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("erro ao testar conexão com banco: %v", err)
	}
}

func SalvarLotes(records [][]any) error {
	chunks := SepararEmLotes(records, batchSize)

	jobs := make(chan [][]any, len(chunks))
	errs := make(chan error, len(chunks))
	var wg sync.WaitGroup

	for w := 0; w < maxWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range jobs {
				if err := InserirEmLote(batch); err != nil {
					errs <- err
				}
			}
		}()
	}

	for _, batch := range chunks {
		jobs <- batch
	}
	close(jobs)

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func InserirEmLote(records [][]any) error {
	if len(records) == 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var b strings.Builder
	b.WriteString(`
		INSERT INTO cliente_compras (
			CPF, IS_CPF_VALIDO, PRIVATE, INCOMPLETO, ULTIMA_COMPRA,
			TICKET_MEDIO, TICKET_ULTIMA_COMPRA, LOJA_FREQUENTE,
			IS_LOJA_FREQUENT_VALIDO, LOJA_ULTIMA_COMPRA, IS_LOJA_ULTIMA_COMPRA_VALIDO
		) VALUES 
	`)

	args := make([]interface{}, 0, len(records)*11)
	argCount := 1

	for i, rec := range records {
		b.WriteString("(")
		for j := 0; j < len(rec); j++ {
			b.WriteString(fmt.Sprintf("$%d", argCount))
			argCount++
			if j < len(rec)-1 {
				b.WriteString(",")
			}
			args = append(args, rec[j])
		}
		b.WriteString(")")
		if i < len(records)-1 {
			b.WriteString(",")
		}
	}

	_, err = tx.Exec(b.String(), args...)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func SepararEmLotes(data [][]any, size int) [][][]any {
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
