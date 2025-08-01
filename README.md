
# Desafio Técnico | BU Sales & Marketing  🚀

Este projeto em Go tem como objetivo demonstrar habilidades com upload de arquivos, processamento ETL e integração com banco de dados PostgreSQL. Também possui documentação interativa via Swagger.

---

##  Estrutura do Projeto

- `handlers/` – Endpoints da API
- `middleware/` – Autenticação e segurança
- `etl/` – Processamento dos arquivos
- `postgres/` – Scripts de inicialização do banco
- `uploads/` – Diretório onde os arquivos são armazenados
- `util/` – Funções auxiliares
- `docs/` – Documentação gerada pelo Swagger

---

##  Requisitos

- [Go 1.20](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

##  Como rodar o projeto

```bash
# Clone o repositório
git clone https://github.com/wsmonteiro10/teste_go.git
cd teste_go

# Suba a aplicação com Docker
docker-compose up --build
```

A API estará disponível em:  
👉 `http://localhost:8080`

A documentação Swagger estará em:  
👉 `http://localhost:8080/swagger/index.html`

---

##  Documentação da API

###  Health Check
- `GET /ping`  
  Retorna `pong` – usado para verificar se o servidor está online.

---

###  ETL Rápido
- `POST /upload`  
  Upload de arquivo `.txt` ou `.csv`, grava os dados no banco Postgres com índices.

- `POST /upload_semindex`  
  Upload de arquivo `.txt` ou `.csv`, grava direto no banco **sem criar índices** (melhor performance para grandes volumes).

---

###  Arquivos
- `GET /listas_arquivos`  
  Retorna a lista de arquivos que foram enviados para o servidor (diretório `/uploads`).

---

##  Autenticação

As rotas de upload estão protegidas. Utilize o botão `Authorize` no Swagger para fornecer um `Client Secret` (se configurado no middleware).

---

