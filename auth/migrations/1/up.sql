CREATE TYPE "encryption" AS ENUM (
    'sha-256',
    'bcrypt'
    );

CREATE TABLE "credentials" (
    "id" uuid primary key,
    "login" text not null,
    "password" text not null,
    "email" text not null,
    "method" encryption not null);

CREATE TABLE "history" (
    "id" uuid primary key,
    "auth" uuid references "credentials" ("id") on delete cascade,
    "ip" text,
    "time" timestamp
);

CREATE TABLE "restore_token" (
    "id" uuid primary key,
    "auth" uuid references "credentials" ("id") on delete cascade,
    "token" text,
    "expires" timestamp
);