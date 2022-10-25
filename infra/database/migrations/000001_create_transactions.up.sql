CREATE TABLE "transactions" (
    "updated_at" timestamptz,
    "id" bigserial NOT NULL,
    "medium" varchar(50) NOT NULL DEFAULT 'app',
    "from_wallet" text,
    "from_wallet_id" bigint,
    "amount" text,
    "service_charge" text,
    "transaction_type" text,
    "payment_for" varchar(30),
    "reference" text,
    "trx_id" varchar(100) NOT NULL UNIQUE,
    "status" varchar(50) NOT NULL,
    "initiated_persona" varchar(50) NOT NULL,
    "batch_id" varchar(100),
    "initiated_at" timestamptz,
    "completed_at" timestamptz DEFAULT null,
    PRIMARY KEY ("id")
)