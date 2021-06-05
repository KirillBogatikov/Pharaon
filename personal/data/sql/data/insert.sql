INSERT INTO "data" ("id", "phone", "photo_id", "birth_date")
    VALUES (:data_id, :phone, :photo_id, :birth_date);
INSERT INTO "name" ("id", "data_id", "first", "last", "patronymic")
    VALUES (:name_id, :data_id, :first, :last, :patronymic);