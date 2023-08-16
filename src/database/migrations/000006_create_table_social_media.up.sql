DROP TABLE IF EXISTS "social_media";
CREATE TABLE "social_media" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(100) NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

ALTER TABLE "complainant" ADD FOREIGN KEY ("social_media_id") REFERENCES "social_media" ("id");
