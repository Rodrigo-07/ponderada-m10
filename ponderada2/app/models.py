sql_upsert_saldo = """
UPDATE clientes
SET saldo = saldo + $1
WHERE id = $2
  AND saldo + $1 >= -limite
RETURNING limite, saldo;
"""

sql_insert_tx = """
INSERT INTO transacoes (cliente_id, valor, tipo, descricao)
VALUES ($1, $2, $3, $4);
"""

sql_select_saldo = """
SELECT limite, saldo FROM clientes WHERE id = $1;
"""

sql_select_ultimas = """
SELECT valor, tipo, descricao, realizada_em
FROM   transacoes
WHERE  cliente_id = $1
ORDER  BY realizada_em DESC
LIMIT 10;
"""

sql_upsert_e_insert = """
WITH saldo_atual AS (
    UPDATE clientes
       SET saldo = saldo + $2        -- delta (+credito | -debito)
     WHERE id = $1
       AND saldo + $2 >= -limite
 RETURNING saldo, limite
)
INSERT INTO transacoes (cliente_id, valor, tipo, descricao)
SELECT $1, $3, $4, $5               -- id, valor, tipo, descricao
  FROM saldo_atual
RETURNING (SELECT limite FROM saldo_atual) AS limite,
          (SELECT saldo  FROM saldo_atual) AS saldo;
"""
