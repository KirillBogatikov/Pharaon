CREATE TABLE "file" (
    "id" uuid primary key,
    "name" text,
    "size" bigint,
    "expires" timestamp
);

CREATE TABLE "action" (
    "id" uuid primary key,
    "file" uuid references "file"("id"),
    "action" action_type,
    "date" timestamp
);