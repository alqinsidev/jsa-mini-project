DROP TABLE IF EXISTS "sub_category";
CREATE TABLE "sub_category" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(100) NOT NULL,
  "main_category_id" UUID,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

ALTER TABLE "aduan" ADD FOREIGN KEY ("sub_category_id") REFERENCES "sub_category" ("id");