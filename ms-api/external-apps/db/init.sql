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

INSERT INTO videos VALUES(1,'user_1','title 1','2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO videos VALUES(2, 'user_2', 'title 2', '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO videos VALUES(3, 'user_3', 'title 3', '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO videos VALUES(4, 'user_4', 'title 4', '2000-01-01 00:00:00', '2000-01-01 00:00:00');


INSERT INTO rates VALUES(1, 'user_1', 1, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO rates VALUES(2, 'user_2', 2, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO rates VALUES(3, 'user_3', 3, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO rates VALUES(4, 'user_4', 4, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');

INSERT INTO views VALUES(1, 'user_1', 1, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO views VALUES(2, 'user_2', 2, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO views VALUES(3, 'user_3', 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO views VALUES(4, 'user_4', 4, '2000-01-01 00:00:00', '2000-01-01 00:00:00');


SELECT setval('videos_id_seq', (SELECT MAX(id) FROM videos));
SELECT setval('rates_id_seq', (SELECT MAX(id) FROM rates));
SELECT setval('views_id_seq', (SELECT MAX(id) FROM views));