CREATE TABLE "name" (
    "id" uuid primary key,
    "first" text,
    "last" text,
    "patronymic" text
);

CREATE TABLE "data" (
    "id" uuid primary key,
    "name" uuid references "name"("id"),
    "phone" text,
    "birth_date" date,
    "photo" uuid
)