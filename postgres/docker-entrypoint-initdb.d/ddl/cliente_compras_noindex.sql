CREATE TABLE cliente_compras_noindex (
    CPF                          VARCHAR(11) NOT NULL,
    IS_CPF_VALIDO                INTEGER,
    PRIVATE                      INTEGER,
    INCOMPLETO                   INTEGER,
    ULTIMA_COMPRA                DATE,
    TICKET_MEDIO                 NUMERIC(10, 2),
    TICKET_ULTIMA_COMPRA         NUMERIC(10, 2),
    LOJA_FREQUENTE               VARCHAR(14),
    IS_LOJA_FREQUENT_VALIDO      INTEGER,
    LOJA_ULTIMA_COMPRA           VARCHAR(14),
    IS_LOJA_ULTIMA_COMPRA_VALIDO INTEGER
);
