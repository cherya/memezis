-- +goose Up
-- +goose StatementBegin
create table if not exists "posts" (
    "id" bigserial primary key,
	"source" text,
	"submitted_by" text,
	"text" text,
    "tags" bigint[] not null default '{}',
	"created_at" timestamp not null default now(),
	"original_created_at" timestamp not null default now(),
	"has_media" bool not null
);
create index tags_index on posts using gin (tags);

create table if not exists "media" (
	"id" bigserial primary key,
	"post_id" bigint not null,
	"url" text,
	"type" text,
    "source_id" text,

	foreign key(post_id) references posts(id)
);

create table if not exists "votes_count" (
	"id" bigserial primary key,
	"post_id" bigint not null,
	"up" int default 0,
	"down" int default 0,

	foreign key(post_id) references posts(id),
	unique(post_id)
);

create table if not exists "votes" (
	"id" bigserial primary key,
	"post_id" bigint not null,
	"user_id" text,
	"is_up" bool,
	"created_at" timestamp not null default now(),

	unique (post_id, user_id)
);

create table if not exists "publish" (
	"id" serial primary key,
	"post_id" bigint not null,
	"status" text,
	"published_at" timestamp,
	"published_to" text,

    foreign key(post_id) references posts(id)
);

create table if not exists "tags" (
	"id" serial primary key,
	"text" text,
    unique(text)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tags;
drop table publish;
drop table votes;
drop table votes_count;
drop table media;
drop table posts;
-- +goose StatementEnd
