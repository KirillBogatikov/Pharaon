SELECT
    "s"."id" AS "id",
    "s"."name" AS "name",
    "s"."version" AS "version"
FROM "service" "s"
WHERE "s"."name" = $1