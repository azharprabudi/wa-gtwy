package wasql

const V1 = `
CREATE TABLE "sessions" (
	"id" uuid NOT NULL,
	"value" text NOT NULL,
	"created_at" timestampz NOT NULL,
	CONSTRAINT "sessions_pk" PRIMARY KEY ("id"),
);
`
