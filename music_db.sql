create database music_service;
CREATE USER andrew WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE music_service TO andrew;

CREATE TABLE IF NOT EXISTS genres
(
    genre_id serial PRIMARY KEY,
    title    varchar(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS musicians
(
    musician_id serial PRIMARY KEY,
    name        varchar(40) NOT NULL,
    description text,
    picture     varchar(100)
);

CREATE TABLE Musicians_to_Genres
(
    genre_id INTEGER NOT NULL,
    musician_id INTEGER NOT NULL,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS albums
(
    album_id     serial PRIMARY KEY,
    tittle       varchar(100),
    picture      varchar(100),
    release_date date
);

CREATE TABLE IF NOT EXISTS tracks
(
    track_id     serial PRIMARY KEY,
    tittle       varchar(100),
    text         text,
    audio        bytea,
    picture      varchar(100),
    release_date date
);

CREATE TABLE Musicians_to_Tracks
(
    track_id    INTEGER NOT NULL,
    musician_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE
);

CREATE TABLE Tracks_to_Genres
(
    track_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE
);

CREATE TABLE Tracks_to_Albums
(
    track_id INTEGER NOT NULL,
    album_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete CASCADE
);

CREATE TABLE Musicians_to_Albums
(
    musician_id INTEGER NOT NULL,
    album_id    INTEGER NOT NULL,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete CASCADE,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE
);

CREATE TABLE Albums_to_Genres
(
    genre_id INTEGER NOT NULL,
    album_id INTEGER NOT NULL,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE
);
