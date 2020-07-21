CREATE TABLE IF NOT EXISTS users
(
	id       INTEGER UNIQUE GENERATED BY DEFAULT AS IDENTITY,
	name     TEXT NOT NULL,
	email    TEXT NOT NULL,
	password TEXT NOT NULL,
	bio      TEXT DEFAULT '',
	image    TEXT DEFAULT '',
	PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS tags
(
	id   INTEGER GENERATED BY DEFAULT AS IDENTITY,
	name TEXT,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS articles
(
	id          INTEGER GENERATED BY DEFAULT AS IDENTITY,
	slug        TEXT,
	title       TEXT NOT NULL,
	description TEXT NOT NULL,
	body        TEXT NOT NULL,
	createdAt   TIMESTAMP DEFAULT now(),
	updatedAt   TIMESTAMP DEFAULT now(),
	author_id   INTEGER,
	PRIMARY KEY (id),
	FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_favorite_articles
(
	id         INTEGER GENERATED BY DEFAULT AS IDENTITY,
	article_id INTEGER,
	user_id    INTEGER,
	PRIMARY KEY (id),
	FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE

);


CREATE TABLE IF NOT EXISTS tag_in_articles
(
	id         INTEGER GENERATED BY DEFAULT AS IDENTITY,
	article_id INTEGER,
	tag_id     INTEGER,
	PRIMARY KEY (id),
	FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE,
	FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
);