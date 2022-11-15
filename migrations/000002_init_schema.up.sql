CREATE TABLE reports (
    id bigserial PRIMARY KEY,
   created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   user_id bigserial NOT NULL,
   post_id bigint NOT NULL
);