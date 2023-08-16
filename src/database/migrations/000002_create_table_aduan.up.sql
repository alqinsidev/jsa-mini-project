DROP TABLE IF EXISTS "aduan";
CREATE TABLE "aduan" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "complain_id" varchar(255) NOT NULL,
  "title" varchar(100) NOT NULL,
  "description" varchar(255) NOT NULL,
  "evidence" text NOT NULL,
  "complainant_id" UUID NOT NULL,
  "complainant_position_lat" varchar(50) NOT NULL,
  "complainant_position_lon" varchar(50) NOT NULL,
  "complained_city" varchar(100) NOT NULL,
  "complained_district" varchar(100) NOT NULL,
  "complained_sub_district" varchar(100) NOT NULL,
  "complained_address" varchar(255) NOT NULL,
  "complained_gmap_link" text NOT NULL,
  "category_id" UUID NOT NULL,
  "sub_category_id" UUID NOT NULL,
  "status" int NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

CREATE INDEX "idx_complain_id" ON "aduan" ("complain_id");