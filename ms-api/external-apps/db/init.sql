-- CREATE DATABASE ms_db;

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
    PRIMARY KEY ("user_id", "video_id")
);

CREATE TABLE views (
    "id" SERIAL UNIQUE NOT NULL,
    "user_id" VARCHAR(255),
    "video_id" INTEGER REFERENCES videos (id) ON DELETE CASCADE,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id")
);

INSERT INTO videos VALUES(1,'user1','title 1','2000-01-01 00:00:00', '2000-01-01 00:00:00')