-- +goose Up
-- +goose StatementBegin
ALTER TABLE "media" ADD COLUMN "sha1" text DEFAULT '';
CREATE INDEX media_sha1_index ON media(sha1);
CREATE INDEX media_source_id_index ON media(source_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "media" DROP COLUMN "sha1";
DROP INDEX media_sha1_index;
DROP INDEX media_source_id_index;
-- +goose StatementEnd
