CREATE TABLE "user" (
    "id" uuid primary key,
    "credentials_id" uuid not null,
    "personal_data_id" uuid not null
);