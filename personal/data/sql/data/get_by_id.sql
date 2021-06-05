SELECT
    "d"."id" AS "data_id",
    "d"."phone" AS "phone",
    "d"."birth_date" AS "birth_date",
    "d"."photo_id" AS "photo_id",
    "n"."id" AS "name_id",
    "n"."first" AS "first",
    "n"."last" AS "last",
    "n"."patronymic" AS "patronymic"
FROM "data" AS "d"
JOIN "name" "n" ON "d"."id" = "n"."data_id"
WHERE "d"."id" = :data_id