ALTER TABLE "data"
    DROP "name";
ALTER TABLE "data"
    RENAME "photo" TO "photo_id";
ALTER TABLE "name"
    ADD "data_id" uuid references "data"("id") on delete cascade;
