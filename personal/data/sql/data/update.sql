UPDATE "data" SET
  "phone" = :phone,
  "photo_id" = :photo_id,
  "birth_date" = :birth_date
WHERE "id" = :data_id;

UPDATE "name" SET
  "first" = :first,
  "last" = :last,
  "patronymic" = :patronymic
WHERE "data_id" = :data_id;