CREATE TABLE "card"
(
    "id"          uuid primary key,
    "name"        text,
    "description" text,
    "author_id"   uuid,
    "executor_id" uuid,
    "type"        card_type,
    "priority"    priority
);

CREATE TABLE "observer"
(
    "id"      uuid primary key,
    "card_id" uuid references "card" ("id") on delete cascade,
    "user_id" uuid
);

CREATE TABLE "tag"
(
    "id"    uuid primary key,
    "label" text
);

CREATE TABLE "card_tag"
(
    "id"      uuid primary key,
    "card_id" uuid references "card" ("id") on delete cascade,
    "tag_id"  uuid references "tag" ("id") on delete cascade
);