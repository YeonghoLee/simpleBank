-- 외래 키 제약 조건 제거
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- users 테이블 삭제
DROP TABLE IF EXISTS "users";

CREATE TABLE "users" (
    "username" VARCHAR PRIMARY KEY,
    "hashed_password" VARCHAR NOT NULL,
    "full_name" VARCHAR NOT NULL,
    "email" VARCHAR UNIQUE NOT NULL,
    "password_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

-- users 테이블에 accounts의 owner 값을 추가
WITH email_assignment AS (
    SELECT DISTINCT 
           "owner", 
           'default_email' || ROW_NUMBER() OVER (ORDER BY "owner") || '@example.com' AS "email"
    FROM "accounts"
    WHERE "owner" IS NOT NULL
    AND NOT EXISTS (SELECT 1 FROM "users" WHERE "username" = "accounts"."owner")
)
INSERT INTO "users" ("username", "hashed_password", "full_name", "email")
SELECT "owner", 
       'default_hashed_password', 
       'default_full_name', 
       "email"
FROM email_assignment
WHERE NOT EXISTS (SELECT 1 FROM "users" WHERE "username" = email_assignment."owner");

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

--CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");