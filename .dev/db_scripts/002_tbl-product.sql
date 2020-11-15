BEGIN;

-- CREATE SEQUENCE "wp_product_id_seq" --------------------------
CREATE SEQUENCE IF NOT EXISTS "public"."wp_product_id_seq"
INCREMENT 1
MINVALUE 1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
-- -------------------------------------------------------------

COMMIT;

BEGIN;

-- CREATE TABLE "wp_product" ------------------------------------
CREATE TABLE IF NOT EXISTS "public"."wp_product" (
    "id" BIGINT DEFAULT nextval('wp_product_id_seq'::regclass) NOT NULL,
    "shop_id" BIGINT  NOT NUll,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "price" BIGINT DEFAULT 0 NOT NULL,
    "status" smallint DEFAULT 1 NOT NULL,
    "created_by" BIGINT DEFAULT 0,
    "created_time" Timestamp Without Time Zone,
    "updated_by" BIGINT DEFAULT 0,
    "updated_time" Timestamp Without Time Zone,
    PRIMARY KEY ( "id" )
);
-- -------------------------------------------------------------

COMMIT;

BEGIN;

INSERT INTO "public"."wp_product" (shop_id,"name",description,price,status,created_time,updated_time)
VALUES
       (1,'Baju Rodaemon','Baju Roda Emon',100000,1,now(),now()),
       (1,'Baju Pikachu','Baju Pika Pika',120000,1,now(),now()),
       (1,'Baju Iguanodon','Baju Dinosaurus',100000,1,now(),now()),
       (1,'Laptop John Thor','Laptop punya si john thor',12100000,1,now(),now());

COMMIT;

BEGIN;

-- CREATE TABLE "wp_product_img" ------------------------------------
CREATE TABLE IF NOT EXISTS "public"."wp_product_img" (
	"product_id" BIGINT  NOT NUll,
	"image_url" VARCHAR(255) NOT NULL,
	"created_by" BIGINT DEFAULT 0,
	"created_time" Timestamp Without Time Zone,
	"updated_by" BIGINT DEFAULT 0,
	"updated_time" Timestamp Without Time Zone
);
-- -------------------------------------------------------------

COMMIT;

BEGIN;

-- CREATE LINK "repository_squad_repository_id_fkey" ----------
ALTER TABLE "public"."wp_product_img"
	ADD CONSTRAINT "wp_product_fkey" FOREIGN KEY ( "product_id" )
	REFERENCES "public"."wp_product"( "id" ) MATCH SIMPLE
	ON DELETE NO ACTION
	ON UPDATE NO ACTION;
-- -------------------------------------------------------------

COMMIT;

BEGIN;

INSERT INTO "public"."wp_product_img" (product_id,image_url,created_time,updated_time)
VALUES
        (1, 'https://via.placeholder.com/400',now(),now()),
        (1, 'https://via.placeholder.com/600',now(),now()),
        (2, 'https://via.placeholder.com/400',now(),now()),
        (2, 'https://via.placeholder.com/600',now(),now()),
        (3, 'https://via.placeholder.com/400',now(),now()),
        (3, 'https://via.placeholder.com/600',now(),now()),
        (4, 'https://via.placeholder.com/400',now(),now()),
        (4, 'https://via.placeholder.com/600',now(),now());

COMMIT;