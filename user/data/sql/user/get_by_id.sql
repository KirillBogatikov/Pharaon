SELECT
    "u"."id" AS "user_id",
    "u"."credentials_id" AS "credentials_id",
    "u"."personal_data_id" AS "personal_data_id"
FROM "user" "u"
WHERE "u"."id" = :user_id