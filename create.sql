DROP database if exists RSS_agregator_telegram_bot;
CREATE database RSS_agregator_telegram_bot;
USE RSS_agregator_telegram_bot;

CREATE TABLE Feeds(
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    url NVARCHAR(1024) NOT NULL UNIQUE,
    standard bool
);

CREATE TABLE Users (
	id INTEGER NOT NULL PRIMARY KEY
);

CREATE TABLE User_Feeds(
	`user`  INTEGER NOT NULL,
    feed  INTEGER NOT NULL,
    description NVARCHAR(1024) DEFAULT NULL,
    PRIMARY KEY(`user`, feed),
    FOREIGN KEY (`user`) REFERENCES Users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (feed) REFERENCES Feeds(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE BotCommands(
  command NCHAR(20) PRIMARY KEY,
  description NVARCHAR(1024) NOT NULL
);

CREATE TABLE SystemInfo(
  lastUpdateId INTEGER PRIMARY KEY
);

INSERT INTO SystemInfo(lastUpdateId) VALUE (0);

INSERT INTO Feeds(url, standard) VALUES
	("http://gazeta.ru/export/rss/lenta.xml", true),
	("http://tvrain.ru/export/rss/programs/1018.xml", true),
	("http://interfax.ru/rss.asp", true),
	("https://buzzfeed.com/index.xml", true),
	("http://feeds.bbci.co.uk/news/world/rss.xml", true),
	("http://news.rambler.ru/rss/world/", true),
	("https://meduza.io/rss/all", true),
	("http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml", true);

INSERT INTO BotCommands VALUES
	(`myfeeds`,    `Show me my feeds`),
	(`addfeed`,    `Add new feed to my feeds`),
	(`removefeed`, `Remove feed from my feeds`);
