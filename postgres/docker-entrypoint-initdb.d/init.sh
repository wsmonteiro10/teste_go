export PGPASSWORD=admin

# Database DDL
psql -U postgres -d postgres -a -f /docker-entrypoint-initdb.d/ddl/cliente_compras.sql

psql -U postgres -d postgres -a -f /docker-entrypoint-initdb.d/ddl/cliente_compras_noindex.sql
