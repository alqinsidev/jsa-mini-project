DROP TABLE IF EXISTS "category";
CREATE TABLE "category" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(100) NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

ALTER TABLE "aduan" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");