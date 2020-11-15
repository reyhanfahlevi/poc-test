BEGIN;

-- CREATE SEQUENCE "wp_shop_id_seq" --------------------------
CREATE SEQUENCE IF NOT EXISTS "public"."wp_shop_id_seq"
INCREMENT 1
MINVALUE 1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
-- -------------------------------------------------------------

COMMIT;

BEGIN;

-- CREATE TABLE "wp_shop" ------------------------------------
CREATE TABLE IF NOT EXISTS "public"."wp_shop" (
	"id" Bigint DEFAULT nextval('wp_shop_id_seq'::regclass) NOT NULL,
	"user_id" Bigint  NOT NUll,
	"name" VARCHAR(255) NOT NULL,
	"address" VARCHAR(255) NOT NULL,
	"status" smallint DEFAULT 1 NOT NULL,
	"created_by" Bigint DEFAULT 0,
	"created_time" Timestamp Without Time Zone,
	"updated_by" Bigint DEFAULT 0,
	"updated_time" Timestamp Without Time Zone,
	PRIMARY KEY ( "id" ),
	CONSTRAINT "unique_user_id" UNIQUE( "user_id" )
);
-- -------------------------------------------------------------

COMMIT;

BEGIN;

INSERT INTO "public"."wp_shop" ("name",user_id,address,status,created_time,updated_time)
VALUES ('Demo Store',1,'Jl. Anggrek No 10, Jakarta',1,now(),now());

COMMIT;
