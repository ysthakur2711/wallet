CREATE TYPE "user_status" AS ENUM (
  'ACTIVE',
  'BLOCKED'
);

CREATE TYPE "wallet_status" AS ENUM (
  'ACTIVE',
  'INACTIVE'
);

CREATE TYPE "payment_request_status" AS ENUM (
  'WAITING_APPROVAL',
  'APPROVED',
  'REFUSED',
  'PAYMENT_SUCCESS',
  'PAYMENT_FAILED'
);

CREATE TABLE "user"
(
    "id"                  bigserial PRIMARY KEY,
    "username"            varchar     NOT NULL,
    "hashed_password"     varchar     NOT NULL,
    "status"              user_status NOT NULL,
    "email"               varchar     NOT NULL,
    "address"             varchar     NOT NULL,
    "nationality"         varchar     NOT NULL,
    "aadhar_no"           varchar     NOT NULL,
    "password_changed_at" timestamp   NOT NULL DEFAULT 'now()',
    "created_at"          timestamp   NOT NULL DEFAULT 'now()',
    "updated_at"          timestamp   NOT NULL DEFAULT 'now()'
);

CREATE TABLE "wallets"
(
    "id"                     bigserial PRIMARY KEY,
    "status"                wallet_status NOT NULL,
    "user_id"                bigint        NOT NULL,
    "balance"                bigint        NOT NULL,
    "currency"               varchar       NOT NULL,
    "created_at"             timestamp     NOT NULL DEFAULT 'now()',
    "updated_at"             timestamp     NOT NULL DEFAULT 'now()'
);

CREATE TABLE "trans"
(
    "trans_key"             bigserial PRIMARY KEY,
    "from_acc_id"           bigint    NOT NULL,
    "to_acc_id"             bigint    NOT NULL,
    "amount"                bigint    NOT NULL,
    "created_at"     timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "wallets"
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "trans"
    ADD FOREIGN KEY ("from_acc_id") REFERENCES "wallets" ("id");