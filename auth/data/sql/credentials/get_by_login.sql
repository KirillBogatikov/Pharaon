SELECT
    "c"."id" as "auth_id",
    "c"."login" as "login",
    "c"."password" as "password",
    "c"."email" as "email",
    "c"."method" as "method"
FROM "credentials" "c"
WHERE "c"."login" = $1