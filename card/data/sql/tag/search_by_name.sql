SELECT
    "t"."id" AS "tag_id",
    "t"."name" AS "tag_name"
FROM "tag" "t"
WHERE "name" ~ :tag_name