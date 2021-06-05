SELECT
   "t"."id" AS "token_id",
   "t"."auth" AS "auth_id",
   "t"."token" AS "token",
   "t"."expires" AS "expires"
FROM "restore_token" "t"
WHERE "t"."token" = $1