UPDATE "card" SET
  "name" = :card_name,
  "description" = :description,
  "author_id" = :author_id,
  "executor_id" = :executor_id,
  "type" = :type,
  "priority" = :priority
WHERE "id" = :card_id
