CREATE USER andrew WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE music_service_docker TO andrew;

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS genres
(
    genre_id serial PRIMARY KEY,
    title    varchar(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS musicians
(
    musician_id serial PRIMARY KEY,
    name        citext NOT NULL,
    description text,
    picture     varchar(100)
);

CREATE TABLE IF NOT EXISTS Musicians_to_Genres
(
    genre_id    INTEGER NOT NULL,
    musician_id INTEGER NOT NULL,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS albums
(
    album_id     serial PRIMARY KEY,
    tittle       citext,
    picture      varchar(100),
    release_date date
);

CREATE TABLE IF NOT EXISTS tracks
(
    track_id          serial PRIMARY KEY,
    tittle            varchar(100),
    text              text,
    rating            int default 0,
    amount_of_listens int default 0,
    audio             bytea,
    picture           varchar(100),
    release_date      date
);

-- ///

CREATE TABLE IF NOT EXISTS playlists
(
    playlist_id  serial PRIMARY KEY,
    user_id      int not null,
    tittle       varchar(100),
    description  text,
    picture      varchar(100),
    release_date date
);

CREATE TABLE IF NOT EXISTS playlists_to_genres
(
    playlist_id INTEGER NOT NULL,
    genre_id    INTEGER NOT NULL,
    FOREIGN KEY (playlist_id) REFERENCES playlists (playlist_id) on delete CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS playlists_to_user
(
    user_id     INTEGER NOT NULL,
    playlist_id INTEGER NOT NULL,
    FOREIGN KEY (playlist_id) REFERENCES playlists (playlist_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Tracks_to_Playlist
(
    track_id    INTEGER NOT NULL,
    playlist_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (playlist_id) REFERENCES playlists (playlist_id) on delete CASCADE
);


CREATE TABLE IF NOT EXISTS Musicians_to_Playlist
(
    musician_id INTEGER NOT NULL,
    playlist_id INTEGER NOT NULL,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE,
    FOREIGN KEY (playlist_id) REFERENCES playlists (playlist_id) on delete CASCADE
);

-- ///

CREATE TABLE IF NOT EXISTS Musicians_to_Tracks
(
    track_id    INTEGER NOT NULL,
    musician_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Tracks_to_Genres
(
    track_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Tracks_to_Albums
(
    track_id INTEGER NOT NULL,
    album_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Musicians_to_Albums
(
    musician_id INTEGER NOT NULL,
    album_id    INTEGER NOT NULL,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete CASCADE,
    FOREIGN KEY (musician_id) REFERENCES Musicians (musician_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Albums_to_Genres
(
    genre_id INTEGER NOT NULL,
    album_id INTEGER NOT NULL,
    FOREIGN KEY (album_id) REFERENCES Albums (album_id) on delete CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Profiles
(
    profiles_id        bigserial not null primary key,
    email              varchar   not null unique,
    nickname           citext    not null unique,
    first_name         varchar   not null,
    second_name        varchar   not null,
    encrypted_password varchar   not null,
    avatar             varchar   not null,
    favorite_genre     text[]    not null default '{}'::text[]
);

INSERT INTO tracks (tittle, text, audio, picture, release_date)
VALUES ('track1', 'some text', 'audio1', 'picture', '2020-03-04'),
       ('track2', 'some text', 'audio2', 'picture', '2020-03-04'),
       ('track3', 'some text', 'audio3', 'picture', '2020-03-04'),
       ('track4', 'some text', 'audio4', 'picture', '2020-03-04');


CREATE TABLE IF NOT EXISTS tracks_to_user
(
    user_id  INTEGER NOT NULL,
    track_id INTEGER NOT NULL,
    favorite bool default false,
    FOREIGN KEY (track_id) REFERENCES tracks (track_id) on delete CASCADE
);


INSERT INTO genres (title)
VALUES ('classical'),
       ('jazz'),
       ('rap'),
       ('electronic'),
       ('rock'),
       ('disco'),
       ('fusion'),
       ('pop'),
       ('country'),
       ('blues'),
       ('reggae'),
       ('indie');