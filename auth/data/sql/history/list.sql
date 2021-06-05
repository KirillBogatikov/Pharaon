SELECT
    "h"."id" AS "history_id",
    "h"."auth" AS "auth_id",
    "h"."ip" AS "ip",
     "h"."time" AS "time"
FROM "history" "h"
WHERE "h"."auth" = $1