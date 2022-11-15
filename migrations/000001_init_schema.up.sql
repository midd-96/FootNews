
CREATE TABLE users (
   id bigserial PRIMARY KEY,
   created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   email TEXT UNIQUE NOT NULL,
   name TEXT NOT NULL,
   password_hash bytea NOT NULL,
   activated bool NOT NULL DEFAULT false
);


CREATE TABLE posts (
   id bigserial PRIMARY KEY,
   created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   title text UNIQUE NOT NULL,
   url text NOT NULL,
   approved bool DEFAULT false,
   user_id bigserial NOT NULL
);

CREATE TABLE comments (
      id bigserial PRIMARY KEY,
      created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
      body text NOT NULL,
      post_id bigint NOT NULL,
      user_id bigserial NOT NULL 
);

CREATE TABLE votes (
   created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   user_id bigserial NOT NULL,
   post_id bigint NOT NULL,
   PRIMARY KEY (user_id, post_id)
);


CREATE TABLE admin (
   id bigserial PRIMARY KEY,
   created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   email TEXT UNIQUE NOT NULL,
   name TEXT NOT NULL,
   password_hash bytea NOT NULL
);


