#  Cria banco de dados e tabelas

DROP DATABASE IF EXISTS pismo;
CREATE DATABASE pismo;
USE pismo;

DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts
(
  Account_id                 INT UNSIGNED   NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Available_credit_limit     FLOAT UNSIGNED NOT NULL,
  Available_withdrawal_limit FLOAT UNSIGNED NOT NULL
);

DROP TABLE IF EXISTS operation_types;
CREATE TABLE IF NOT EXISTS operation_types
(
  Operation_type_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Description       VARCHAR(255) NOT NULL,
  Charge_order      SMALLINT     NOT NULL
);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions
(
  Transaction_id    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  Account_id        INT UNSIGNED    NOT NULL,
  Operation_type_id INT UNSIGNED    NOT NULL,
  Amount            FLOAT           NOT NULL CHECK ( Amount < 0 ),
  Balance           FLOAT           NOT NULL CHECK ( Balance <= 0 ),
  Created_at        TIMESTAMP,
  Due_at            DATETIME,
  CONSTRAINT FK_account_transaction
    FOREIGN KEY (Account_id)
      REFERENCES accounts (Account_id)
      ON UPDATE CASCADE
      ON DELETE RESTRICT,
  CONSTRAINT FK_operation_type_transaction
    FOREIGN KEY (Operation_type_id)
      REFERENCES operation_types (Operation_type_id)
      ON UPDATE CASCADE
      ON DELETE RESTRICT
);

INSERT INTO accounts
VALUES (1, 5000, 5000);

INSERT INTO operation_types
VALUES (1, 'COMPRA A VISTA', 2);
INSERT INTO operation_types
VALUES (2, 'COMPRA PARCELADA', 1);
INSERT INTO operation_types
VALUES (3, 'SAQUE', 0);
INSERT INTO operation_types
VALUES (4, 'PAGAMENTO', 0);

INSERT INTO transactions
VALUES (1, 1, 1, -50, -50, '2017-04-05', '2017-05-10');
INSERT INTO transactions
VALUES (2, 1, 1, -23.5, -23.5, '2017-04-10', '2017-05-10');
INSERT INTO transactions
VALUES (3, 1, 1, -18.7, -18.7, '2017-04-30', '2017-06-10');
