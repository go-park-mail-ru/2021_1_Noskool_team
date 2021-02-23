CREATE TABLE IF NOT EXISTS Genres (
                                      genre_id serial PRIMARY KEY,
                                      title varchar(30) NOT NULL
    );

CREATE TABLE IF NOT EXISTS Musicians (
                                         musician_id serial PRIMARY KEY,
                                         name varchar(40) NOT NULL,
    description text,
    picture varchar(100),
    genre_id int,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete SET DEFAULT
    );

CREATE TABLE IF NOT EXISTS Albums (
                                      album_id serial PRIMARY KEY,
                                      tittle varchar(100),
    picture varchar(100),
    release_date date,
    genre_id int,
    musician_id int,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete SET DEFAULT,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete SET DEFAULT
    );

CREATE TABLE IF NOT EXISTS Songs (
                                     song_id serial PRIMARY KEY,
                                     tittle varchar(100),
    text text,
    audio bytea,
    picture varchar(100),
    release_date date,
    genre_id int,
    musician_id int,
    album_id int,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete SET DEFAULT,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete SET DEFAULT,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete SET DEFAULT
    );

