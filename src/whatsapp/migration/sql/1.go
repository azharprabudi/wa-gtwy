package wasql

const V1 = `
CREATE TABLE "sessions" ( "id" uuid not null, "value" text not null, "created_at" timestamp not null, constraint "sessions_pk" primary key ("id") );
`
