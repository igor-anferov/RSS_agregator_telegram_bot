DROP database if exists RSS_agregator_telegram_bot;
CREATE database RSS_agregator_telegram_bot;
USE RSS_agregator_telegram_bot;

CREATE TABLE Feeds(
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    url NVARCHAR(1024) NOT NULL UNIQUE,
    standard bool
);

CREATE TABLE Users (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE User_Feeds(
	`user`  INTEGER NOT NULL REFERENCES Users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    feed  INTEGER NOT NULL REFERENCES Feeds(id) ON DELETE CASCADE ON UPDATE CASCADE,
    description NVARCHAR(1024) NOT NULL UNIQUE,
    PRIMARY KEY(`user`, feed)
);
