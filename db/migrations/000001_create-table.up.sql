CREATE TABLE "transactions" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "from_account_id" integer NOT NULL,
  "to_account_id" integer NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar NOT NULL,
  "ammount" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" integer PRIMARY KEY,
  "amount" varchar NOT NULL,
  "account_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX "from_account_id" ON "transactions" ("from_account_id");

CREATE INDEX "to_account_id" ON "transactions" ("to_account_id");

CREATE INDEX "from_to_account_id" ON "transactions" ("from_account_id", "to_account_id");

CREATE INDEX "account_username" ON "accounts" ("username");

CREATE INDEX "entries_account_id" ON "entries" ("account_id");

COMMENT ON COLUMN "transactions"."amount" IS 'Have be positive';

COMMENT ON COLUMN "entries"."amount" IS 'can be negative';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
