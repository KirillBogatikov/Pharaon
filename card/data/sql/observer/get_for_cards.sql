SELECT
    "o"."id" AS "observer_id",
    "o"."card_id" AS "card_id",
    "o"."user_id" AS "user_id"
FROM "observer" "o"
WHERE "card_id" IN :card_id_list