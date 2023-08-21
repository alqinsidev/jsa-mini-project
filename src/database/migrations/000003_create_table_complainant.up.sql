DROP TABLE IF EXISTS "complainant";
CREATE TABLE "complainant" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(100) NOT NULL,
  "phone" varchar(20) NOT NULL,
  "email" varchar(100) NOT NULL,
  "social_media_id" UUID NOT NULL,
  "social_media_link" text NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

CREATE INDEX "idx_complainant_name" ON "complainant" ("name");

ALTER TABLE "aduan" ADD FOREIGN KEY ("complainant_id") REFERENCES "complainant" ("id");