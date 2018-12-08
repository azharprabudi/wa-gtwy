package wasql

const V1 = `
CREATE TABLE "user_sessions" (
	"id" uuid NOT NULL,
	"value" text NOT NULL,
	"created_at" timestampz NOT NULL,
	"expired_at" timestampz, 
	CONSTRAINT "user_sessions_pk" PRIMARY KEY ("id"),
);
`
