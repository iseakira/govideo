-- CREATE DATABASE ms\db

\c ms_db

CREATE TABLE videos (
  "id" SERIAL UNIQUE NOT NULL,
  "user_id" VARCHAR(255),
  "title" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE rates (
  "id" SERIAL UNIQUE NOT NULL,
  "user_id" VARCHAR(255),
  "video_id" INTEGER REFERENCES videos (id) ON DELETE CASCADE,
  "value" DECIMAL DEFAULT 3.0,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,
  PRIMARY KEY ("user_id","video_id")
);

CREATE TABLE views (
  "id" SERIAL UNIQUE NOT NULL,
  "user_id" VARCHAR(255),
  "video_id" INTEGER REFERENCES videos (id) ON DELETE CASCADE,
  "value" DECIMAL DEFAULT 3.0,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL,
  PRIMARY KEY ("id")
);

CREATE INDEX ON rates(video_id)
CREATE INDEX ON views(video_id)
CREATE INDEX ON videos(user_id)