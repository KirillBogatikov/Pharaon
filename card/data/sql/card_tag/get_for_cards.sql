SELECT
    "ct"."id" AS "cart_tag_id",
    "ct"."card_id" AS "card_id",
    "t"."id" AS "tag_id",
    "t"."name" AS "tag_name"
FROM "card_tag" AS "ct"
JOIN "tag" "t" ON "t"."id" = "ct"."tag_id"
WHERE "ct"."card_id" IN :card_id_list