CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "account_name" varchar NOT NULL,
  "balance" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "sender_id" bigint,
  "receiver_id" bigint,
  "amount" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "account" ("id");

CREATE INDEX ON "transfers" ("sender_id");

CREATE INDEX ON "transfers" ("receiver_id");

CREATE INDEX ON "transfers" ("sender_id", "receiver_id");

CREATE INDEX ON "entries" ("account_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("sender_id") REFERENCES "account" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("receiver_id") REFERENCES "account" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");
