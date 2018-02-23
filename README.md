FN
==

FetchNews - news aggregator.

Service to fetch news from websites or RSS feeds.

User may specify news sources to monitor.

FetchNews will monitor these sources and save news articles into local database (SQLite).

Later user may ask FetchNews to show collected articles (optionally filtered by title substring)

Software design
---------------

Application is designed with [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) in mind.

Base things are defined in package _news_.

Article repository is implemented in package _news/articlerepo/sqlite_.

News source is implemented in packages _news/source/http_ and _news/source/rss_.

UI is implemented in package _news/ui/telegram_.

Main() initializes article repository and starts ui. UI receives user commands and creates news source objects if needed.

Build and run
-------------

(not tested build instructions myself yet)

Provided that:

* you have installed Docker and Go and go dep
* you have telegram bot and token

```
$ go get github.com/gorilych/fn/
$ cd $GOPATH/src/github.com/gorilych/fn/
$ dep ensure
$ build main.go
$ docker build -t <imagetag> .
```

For pure docker:

```
$ docker run --rm -i -t -e "TGBOTTOKEN=<your telegram bot token here>" <imagetag>
```

For kubernetes, create secret _tg-fetchnews-bot_ with key _token_ according to [documentation](https://kubernetes.io/docs/concepts/configuration/secret/#creating-a-secret-manually) and create deployment with (you might want to replace image tag in yaml beforehand)

```
$ kubectl create -f k8s.fetchnewsbot.yaml
```

Next, connect to your telegram bot. Start with /help command.

Known issues
------------

/stop command may be fetched next time bot runs.

