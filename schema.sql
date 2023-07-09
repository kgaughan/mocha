CREATE TABLE users (
    id       INTEGER NOT NULL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE UNIQUE INDEX username ON users (username);

CREATE TABLE feeds (
    id            INTEGER NOT NULL PRIMARY KEY,
    uuid          TEXT NOT NULL,
    title         TEXT NOT NULL,
    feed_url      TEXT NOT NULL UNIQUE,
    site_url      TEXT NOT NULL,
    checked_at    TEXT DEFAULT '', -- DATETIME, UTC
    etag          TEXT DEFAULT '',
    last_modified TEXT DEFAULT ''  -- DATETIME, UTC
);

CREATE UNIQUE INDEX feed_uuid ON feeds (uuid);
CREATE UNIQUE INDEX feed_url ON feeds (feed_url);

CREATE TABLE entries (
    id            INTEGER NOT NULL PRIMARY KEY,
    feed_id       INTEGER NOT NULL,
    uuid          TEXT NOT NULL UNIQUE,
    url           TEXT NOT NULL,
    title         TEXT NOT NULL,
    summary       TEXT,
    content       TEXT,
    author        TEXT,
    published     TEXT NOT NULL, -- DATETIME, UTC
    updated       TEXT NOT NULL, -- DATETIME, UTC

    FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE
);

CREATE TABLE categories (
    id       INTEGER NOT NULL PRIMARY KEY,
    user_id  INTEGER NOT NULL,
    category TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE subscriptions (
    feed_id     INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,
    category_id INTEGER NOT NULL,

    PRIMARY KEY (feed_id, user_id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
) WITHOUT ROWID;

CREATE TABLE links (
    id       INTEGER NOT NULL PRIMARY KEY,
    entry_id INTEGER NOT NULL,
    ref      TEXT NOT NULL, -- alternate, related, enclosure, via (self is ignored)
    title    TEXT NOT NULL,
    url      TEXT NOT NULL,
    size     INTEGER NOT NULL DEFAULT 0,
    mimetype TEXT,

    FOREIGN KEY (entry_id) REFERENCES entries (id) ON DELETE CASCADE
);

CREATE TABLE entry_states (
    entry_id INTEGER NOT NULL,
    user_id  INTEGER NOT NULL,
    is_read  BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY (entry_id, user_id)
) WITHOUT ROWID;
