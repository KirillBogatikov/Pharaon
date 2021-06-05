ALTER TABLE "data"
    ADD "name" uuid references "name"("id");
ALTER TABLE "data"
    RENAME "photo_id" TO "photo";
ALTER TABLE "name"
    DROP "data_id";
