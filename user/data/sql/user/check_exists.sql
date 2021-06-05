SELECT
    "u"."id" AS "id"
FROM "user" "u"
WHERE "u"."id" IN :user_ids
