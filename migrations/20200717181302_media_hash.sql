-- +goose Up
-- +goose StatementBegin
ALTER TABLE "media" DROP COLUMN IF EXISTS "sha1";
DROP INDEX IF EXISTS media_sha1_index;

ALTER TABLE "media" ADD COLUMN "phash" text DEFAULT '';
CREATE INDEX IF NOT EXISTS media_phash_index ON media(phash);

ALTER TABLE "publish" ADD COLUMN "url" text DEFAULT '';

ALTER TABLE "posts" ADD COLUMN "source_url" text DEFAULT '';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "media" ADD COLUMN "sha1" text DEFAULT '';
CREATE INDEX IF NOT EXISTS media_sha1_index ON media(sha1);

ALTER TABLE "media" DROP COLUMN "phash";
DROP INDEX IF EXISTS media_phash_index;

ALTER TABLE "publish" DROP COLUMN "url";

ALTER TABLE "posts" DROP COLUMN "source_url";
-- +goose StatementEnd
