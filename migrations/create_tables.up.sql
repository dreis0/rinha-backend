CREATE UNLOGGED TABLE customer
(
    id              INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_limit   INT       NOT NULL,
    initial_balance INT       NOT NULL DEFAULT 0,
    created_at      TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNLOGGED TABLE transaction
(
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id INT       NOT NULL,
    value       INT       NOT NULL,
    type        CHAR      NOT NULL,
    description VARCHAR   NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (customer_id) REFERENCES customer (id)
);

CREATE UNLOGGED TABLE customer_balance
(
    customer_id INT       NOT NULL PRIMARY KEY,
    balance     INT       NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (customer_id) REFERENCES customer (id)
);