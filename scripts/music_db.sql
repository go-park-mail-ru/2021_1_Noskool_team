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
    picture     varchar(100),
    rating      int default 0
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

insert into tracks (track_id, tittle, text, audio, picture, release_date)
VALUES (1, 'Do I Wanna Know?', 'text', '/api/v1/data/audio/Do_I_Wanna_Know.ogg', '/api/v1/data/img/tracks/AM.jpeg', '2013-03-03'),
       (2, 'R U Mine', 'some text', '/api/v1/data/audio/R_U_Mine.ogg', '/api/v1/data/img/tracks/AM.jpeg', '2013-03-03'),
       (3, 'One For The Road', 'some text', '/api/v1/data/audio/One_For_The_Road.ogg', '/api/v1/data/img/tracks/AM.jpeg', '2013-03-03'),
       (4, 'Arabella', 'some text', '/api/v1/data/audio/Arabella.ogg', '/api/v1/data/img/tracks/AM.jpeg', '2013-03-03'),
       (5, 'I Want It All', 'some text', '/api/v1/data/audio/I_Want_It_All.ogg', '/api/v1/data/img/tracks/AM.jpeg', '2013-03-03'),
       (6, 'Pretty Boy', 'some text', '/api/v1/data/audio/Joji_feat._Lil_Yachty_Pretty_Boy.ogg', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (7, 'Tick Tock', 'some text', '/api/v1/data/audio/Joji_-_Tick_Tock.ogg', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (8, 'Daylight', 'some text', '/api/v1/data/audio/Joji__Diplo_-_Daylight.ogg', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (9, 'Upgrade', 'some text', '/api/v1/data/audio/Joji_-_Upgrade.ogg', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (10, 'Mr. Hollywood', 'some text', '/api/v1/data/audio/Joji_-_Mr._Hollywood.ogg', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (11, 'Run', 'some text', '/api/v1/data/audio/Joji_-_Run.ogg', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (12, 'Flowers', 'some text', '/api/v1/data/audio/Flowers.ogg', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-10-01'),
       (13, 'Scary Love', 'some text', '/api/v1/data/audio/Scary_Love.ogg', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-10-01'),
       (14, 'Nervous', 'some text', '/api/v1/data/audio/Nervous.ogg', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-10-01'),
       (15, 'Void', 'some text', '/api/v1/data/audio/Void.ogg', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-10-01'),
       (16, 'Softcore', 'some text', '/api/v1/data/audio/Softcore.ogg', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-10-01'),
       (17, 'Blue', 'some text', '/api/v1/data/audio/Blue.ogg', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-10-01'),
       (18, 'Smells Like Teen Spirit', 'some text', '/api/v1/data/audio/Smells_Like_Teen_Spirit.ogg', '/api/v1/data/img/tracks/smels_like.jpg', '1991-10-01');



insert into albums (album_id, tittle, picture, release_date)
values (1, 'AM', '/api/v1/data/img/tracks/AM.jpeg', '2013-03-03'),
       (2, 'Nectar', '/api/v1/data/img/tracks/Nectar.png', '2018-10-01'),
       (3, 'The Neighbourhood', '/api/v1/data/img/tracks/The_Neighbourhood.jpg', '2018-06-01');

insert into tracks_to_albums (track_id, album_id)
values (1, 1), (2, 1), (3, 1), (4, 1), (5, 1), (6, 2), (7, 2), (8, 2), (9, 2),
       (10, 2), (11, 2), (12, 3), (13, 3), (14, 3), (15, 3), (16, 3), (17, 3);

insert into musicians (musician_id, name, description, picture)
values (1, 'Arctic Monkeys', 'british alternaitve group', '/api/v1/data/img/musicians/arctics_monkeys.jpeg'),
       (2, 'Joji', 'Джордж Кусуноки Миллер, более известный по сценическому псевдониму Joji',
        '/api/v1/data/img/musicians/joji.jpeg'),
       (3, 'The Neighbourhood', 'alternative group', '/api/v1/data/img/musicians/the_neighbourhood.jpeg'),
       (4, 'Nirvana', 'grange', '/api/v1/data/img/musicians/Nirvana.jpeg');

insert into musicians_to_tracks (track_id, musician_id)
values (1, 1), (2, 1), (3, 1), (4, 1), (5, 1), (6, 2), (7, 2), (8, 2), (9, 2), (10, 2), (11, 2),
       (12, 3), (13, 3), (14, 3), (15, 3), (16, 3), (17, 3), (18, 4);

insert into musicians_to_albums (musician_id, album_id)
values (1, 1), (2, 2), (3, 3);

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