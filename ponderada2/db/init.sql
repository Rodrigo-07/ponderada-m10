-- DDL
CREATE TABLE clientes (
    id      INT PRIMARY KEY,
    limite  INT NOT NULL,
    saldo   INT NOT NULL DEFAULT 0
);

CREATE TABLE transacoes (
    id            BIGSERIAL PRIMARY KEY,
    cliente_id    INT NOT NULL REFERENCES clientes(id),
    valor         INT NOT NULL,
    tipo          CHAR(1) NOT NULL,
    descricao     VARCHAR(10) NOT NULL,
    realizada_em  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5 clientes exigidos
INSERT INTO clientes (id, limite, saldo) VALUES
 (1, 100000, 0),
 (2,  80000, 0),
 (3,1000000, 0),
 (4,10000000,0),
 (5, 500000, 0);
