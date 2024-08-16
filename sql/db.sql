CREATE TABLE tasks(
    uuid VARCHAR(128) NOT NULL PRIMARY KEY,
    progress VARCHAR(128) NOT NULL,
    faces_number SMALLINT DEFAULT -1,
    human_number SMALLINT DEFAULT -1,
    male_mean_age SMALLINT  DEFAULT -1,
    female_mean_age SMALLINT DEFAULT -1
);

CREATE TABLE images( 
    id SERIAL PRIMARY KEY,
    task_id  VARCHAR(128) REFERENCES tasks(uuid),
    title VARCHAR(256) NOT NULL
);

CREATE TABLE faces(
    id SERIAL NOT NULL,
    x INTEGER NOT NULL,
    y INTEGER NOT NULL,
    width INTEGER NOT NULL,
    heigth INTEGER NOT NULL,
    sex varchar(16) NOT NULL,
    age INTEGER NOT NULL,
    image_id INTEGER REFERENCES images(id)
);