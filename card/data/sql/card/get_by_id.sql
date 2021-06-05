SELECT
    "c"."id" AS "card_id",
    "c"."name" AS "card_name",
    "c"."description" AS "description",
    "c"."author_id" AS "author_id",
    "c"."executor_id" AS "executor_id",
    "c"."type" AS "type",
    "c"."priority" AS "priority"
FROM "card" "c"
WHERE "id" = :card_id