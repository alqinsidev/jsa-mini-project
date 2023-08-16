DROP TABLE IF EXISTS "status_history";
CREATE TABLE "status_history" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "aduan_id" UUID,
  "status" int,
  "notes" varchar(255),
  "created_at" timestamptz DEFAULT now()
);

ALTER TABLE "status_history" ADD FOREIGN KEY ("aduan_id") REFERENCES "aduan" ("id");