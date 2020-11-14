BEGIN;

-- CREATE SEQUENCE "wp_user_id_seq" --------------------------
CREATE SEQUENCE IF NOT EXISTS "public"."wp_user_id_seq"
INCREMENT 1
MINVALUE 1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
-- -------------------------------------------------------------

COMMIT;

BEGIN;

-- CREATE TABLE "wp_user" ------------------------------------
CREATE TABLE IF NOT EXISTS "public"."wp_user" (
	"id" Bigint DEFAULT nextval('wp_user_id_seq'::regclass) NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"pwd_hash" VARCHAR(255) NOT NULL,
	"reset_token" VARCHAR(255) NOT NULL,
	"status" smallint DEFAULT 1 NOT NULL,
	"created_by" Bigint DEFAULT 0,
	"created_time" Timestamp Without Time Zone,
	"updated_by" Bigint DEFAULT 0,
	"updated_time" Timestamp Without Time Zone,
	PRIMARY KEY ( "id" ),
	CONSTRAINT "unique_email" UNIQUE( "email" )
);
-- -------------------------------------------------------------

COMMIT;

BEGIN;

INSERT INTO "public"."wp_user" ("name",email,pwd_hash,reset_token,status,created_time,updated_time)
VALUES ('admin','demo@wp.com','3cf2df1c4f6c5e95ce13b4a667e15b5bfceed8e299ed6978cebcf0fa403abc78','',1,now(),now());

COMMIT;
