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
    release_date date,
    rating       int default 0
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
    release_date      date,
    duration          text,
    likes             int default 0
);

-- ///

CREATE TABLE IF NOT EXISTS playlists
(
    playlist_id  serial PRIMARY KEY,
    user_id      int not null,
    tittle       varchar(100),
    description  text,
    picture      varchar(100),
    release_date date default now(),
    rating       int  default 0,
    uid          text default ''
);

CREATE TABLE IF NOT EXISTS album_to_user
(
    user_id  INTEGER NOT NULL,
    album_id INTEGER NOT NULL,
    favorite bool default false,
    UNIQUE (user_id, album_id),
    FOREIGN KEY (album_id) REFERENCES albums (album_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS playlists_to_genres
(
    playlist_id INTEGER NOT NULL,
    genre_id    INTEGER NOT NULL,
    FOREIGN KEY (playlist_id) REFERENCES playlists (playlist_id) on delete CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres (genre_id) on delete CASCADE
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

CREATE TABLE IF NOT EXISTS musicians_to_user
(
    user_id     INTEGER NOT NULL,
    musician_id INTEGER NOT NULL,
    favorite    bool default false,
    UNIQUE (user_id, musician_id),
    FOREIGN KEY (musician_id) REFERENCES musicians (musician_id) on delete CASCADE
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

CREATE TABLE if not exists history
(
    user_id       INTEGER NOT NULL,
    track_id      INTEGER NOT NULL,
    creation_date timestamp default now(),
    FOREIGN KEY (track_id) REFERENCES Tracks (track_id) on delete CASCADE
);

CREATE TABLE IF NOT EXISTS Profiles
(
    profiles_id        bigserial not null primary key,
    email              varchar   not null unique,
    nickname           citext    not null unique,
    first_name         varchar   not null default '',
    second_name        varchar   not null default '',
    encrypted_password varchar   not null,
    avatar             varchar   not null,
    favorite_genre     text[]    not null default '{}'::text[]
);

create table if not exists subscriptions
(
    who     int,
    on_whom int,
    FOREIGN KEY (who) REFERENCES profiles (profiles_id) on delete CASCADE,
    FOREIGN KEY (on_whom) REFERENCES profiles (profiles_id) on delete CASCADE,
    unique (who, on_whom)
);

insert into tracks (track_id, tittle, text, audio, picture, release_date, duration, rating)
VALUES (1, 'Do I Wanna Know?', 'text', '/api/v1/music/data/audio/Do_I_Wanna_Know.ogg',
        '/api/v1/music/data/img/tracks/AM.webp',
        '2013-03-03', '4:32', 0),
       (2, 'R U Mine', 'some text', '/api/v1/music/data/audio/R_U_Mine.ogg', '/api/v1/music/data/img/tracks/AM.webp',
        '2013-03-03', '3:22', 50),
       (3, 'One For The Road', 'some text', '/api/v1/music/data/audio/One_For_The_Road.ogg',
        '/api/v1/music/data/img/tracks/AM.webp', '2013-03-03', '3:43', 100),
       (4, 'Arabella', 'some text', '/api/v1/music/data/audio/Arabella.ogg', '/api/v1/music/data/img/tracks/AM.webp',
        '2013-03-03',
        '3:27', 12),
       (5, 'I Want It All', 'some text', '/api/v1/music/data/audio/I_Want_It_All.ogg',
        '/api/v1/music/data/img/tracks/AM.webp',
        '2013-03-03', '3:04', 200),
       (6, 'Pretty Boy', 'some text', '/api/v1/music/data/audio/Joji_feat._Lil_Yachty_Pretty_Boy.ogg',
        '/api/v1/music/data/img/tracks/Nectar.webp', '2018-10-01', '2:37', 5000),
       (7, 'Tick Tock', 'some text', '/api/v1/music/data/audio/Joji_-_Tick_Tock.ogg',
        '/api/v1/music/data/img/tracks/Nectar.webp',
        '2018-10-01', '2:12', 70),
       (8, 'Daylight', 'some text', '/api/v1/music/data/audio/Joji__Diplo_-_Daylight.ogg',
        '/api/v1/music/data/img/tracks/Nectar.webp', '2018-10-01', '2:44', 20),
       (9, 'Upgrade', 'some text', '/api/v1/music/data/audio/Joji_-_Upgrade.ogg',
        '/api/v1/music/data/img/tracks/Nectar.webp',
        '2018-10-01', '1:30', 200),
       (10, 'Mr. Hollywood', 'some text', '/api/v1/music/data/audio/Joji_-_Mr._Hollywood.ogg',
        '/api/v1/music/data/img/tracks/Nectar.webp', '2018-10-01', '3:22', 3),
       (11, 'Run', 'some text', '/api/v1/music/data/audio/Joji_-_Run.ogg', '/api/v1/music/data/img/tracks/Nectar.webp',
        '2018-10-01', '3:15', 3),
       (12, 'Flowers', 'some text', '/api/v1/music/data/audio/Flowers.ogg',
        '/api/v1/music/data/img/tracks/The_Neighbourhood.webp',
        '2018-10-01', '3:18', 300),
       (13, 'Scary Love', 'some text', '/api/v1/music/data/audio/Scary_Love.ogg',
        '/api/v1/music/data/img/tracks/The_Neighbourhood.webp', '2018-10-01', '3:43', 250),
       (14, 'Nervous', 'some text', '/api/v1/music/data/audio/Nervous.ogg',
        '/api/v1/music/data/img/tracks/The_Neighbourhood.webp',
        '2018-10-01', '4:05', 30),
       (15, 'Void', 'some text', '/api/v1/music/data/audio/Void.ogg',
        '/api/v1/music/data/img/tracks/The_Neighbourhood.webp',
        '2018-10-01', '3:24', 40),
       (16, 'Softcore', 'some text', '/api/v1/music/data/audio/Softcore.ogg',
        '/api/v1/music/data/img/tracks/The_Neighbourhood.webp', '2018-10-01', '3:26', 49),
       (17, 'Blue', 'some text', '/api/v1/music/data/audio/Blue.ogg',
        '/api/v1/music/data/img/tracks/The_Neighbourhood.webp',
        '2018-10-01', '3:10', 42),
       (18, 'Smells Like Teen Spirit', 'some text', '/api/v1/music/data/audio/Smells_Like_Teen_Spirit.ogg',
        '/api/v1/music/data/img/tracks/smels_like.webp', '1991-10-01', '5:01', 500),
       (19, 'Good News', 'some text', '/api/v1/music/data/audio/Mac_Miller.ogg',
        '/api/v1/music/data/img/tracks/good_news.webp',
        '2020-01-01', '5:42', 600),
       (20, 'Каждый раз', 'some text', '/api/v1/music/data/audio/mon_every_time.ogg',
        '/api/v1/music/data/img/tracks/monetohka.webp', '2020-01-01', '3:28', 700),
       (21, 'Born To Die', 'some text', '/api/v1/music/data/audio/Lana_Del_Rey_-_Born_To_Die.ogg',
        '/api/v1/music/data/img/tracks/BornToDie.webp', '2012-01-01', '4:45', 1000),
       (22, 'Dark Paradise', 'some text', '/api/v1/music/data/audio/Lana_Del_Rey_-_Dark_Paradise.ogg',
        '/api/v1/music/data/img/tracks/BornToDie.webp', '2012-01-01', '4:03', 160);


insert into albums (album_id, tittle, picture, release_date)
values (1, 'AM', '/api/v1/music/data/img/tracks/AM.webp', '2013-03-03'),
       (2, 'Nectar', '/api/v1/music/data/img/tracks/Nectar.webp', '2018-10-01'),
       (3, 'The Neighbourhood', '/api/v1/music/data/img/tracks/The_Neighbourhood.webp', '2018-06-01'),
       (4, 'Smells Like Teen Spirit', '/api/v1/music/data/img/tracks/smels_like.webp', '2018-06-01'),
       (5, 'Good News', '/api/v1/music/data/img/tracks/good_news.webp', '2018-06-01'),
       (6, 'Каждый раз', '/api/v1/music/data/img/tracks/monetohka.webp', '2018-06-01'),
       (7, 'Born To Die', '/api/v1/music/data/img/tracks/BornToDie.webp', '2012-06-01');



insert into tracks_to_albums (track_id, album_id)
values (1, 1),
       (2, 1),
       (3, 1),
       (4, 1),
       (5, 1),
       (6, 2),
       (7, 2),
       (8, 2),
       (9, 2),
       (10, 2),
       (11, 2),
       (12, 3),
       (13, 3),
       (14, 3),
       (15, 3),
       (16, 3),
       (17, 3),
       (18, 4),
       (19, 5),
       (20, 6),
       (21, 7),
       (22, 7);

insert into musicians (musician_id, name, description, picture)
values (1, 'Arctic Monkeys', 'british alternaitve group', '/api/v1/music/data/img/musicians/arctics_monkeys.webp'),
       (2, 'Joji', 'Джордж Кусуноки Миллер, более известный по сценическому псевдониму Joji',
        '/api/v1/music/data/img/musicians/joji.webp'),
       (3, 'The Neighbourhood', 'alternative group', '/api/v1/music/data/img/musicians/the_neighbourhood.webp'),
       (4, 'Nirvana', 'grange', '/api/v1/music/data/img/musicians/Nirvana.webp'),
       (5, 'Mac Miller', 'some artist', '/api/v1/music/data/img/tracks/good_news.webp'),
       (6, 'Монеточка', 'some artist', '/api/v1/music/data/img/tracks/monetohka.webp'),
       (7, 'Lana Del Rey', 'some artist', '/api/v1/music/data/img/tracks/BornToDie.webp');

insert into musicians_to_tracks (track_id, musician_id)
values (1, 1),
       (2, 1),
       (3, 1),
       (4, 1),
       (5, 1),
       (6, 2),
       (7, 2),
       (8, 2),
       (9, 2),
       (10, 2),
       (11, 2),
       (12, 3),
       (13, 3),
       (14, 3),
       (15, 3),
       (16, 3),
       (17, 3),
       (18, 4),
       (19, 5),
       (20, 6),
       (21, 7),
       (22, 7);

insert into musicians_to_albums (musician_id, album_id)
values (1, 1),
       (2, 2),
       (3, 3),
       (4, 4),
       (5, 5),
       (6, 6),
       (7, 7);


-- insert into playlists (playlist_id, user_id, tittle, description, picture, release_date)
-- VALUES (1, 0, 'Alternative', 'Alternative music', '/api/v1/music/data/img/playlists/alternative.webp', '2021-06-01'),
--        (2, 0, 'For Good Mood', 'favorite music', '/api/v1/music/data/img/playlists/happy.webp', '2021-07-01');

insert into tracks_to_playlist(track_id, playlist_id)
VALUES (1, 1),
       (11, 1),
       (15, 1),
       (18, 1),
       (22, 2),
       (7, 2),
       (20, 2);

CREATE TABLE IF NOT EXISTS tracks_to_user
(
    user_id  INTEGER NOT NULL,
    track_id INTEGER NOT NULL,
    favorite bool default false,
    UNIQUE (user_id, track_id),
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


insert into Musicians_to_Genres (genre_id, musician_id)
values (1, 4),
       (2, 4),
       (3, 4);