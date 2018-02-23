// UI package for telegram bot

package telegram

import (
	news "github.com/gorilych/fn/news"
	//aRepo "github.com/gorilych/fn/news/articlerepo/sqlite"
	httpSrc "github.com/gorilych/fn/news/source/http"
	rssSrc "github.com/gorilych/fn/news/source/rss"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
  shellwords "github.com/mattn/go-shellwords"
	"time"
  "log"
  "fmt"
)

func debug(msg string) {
  log.Println("UI: " + msg)
}

func Run(repo news.ArticleRepo, token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
    debug("Failed to connect to TG bot API with provided token")
		return err
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	newsSrv := news.NewsService(repo, 10*time.Second)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		reply := "no valid command received"
		if update.Message.IsCommand() {
      debug("Command received: " + update.Message.Command())
			switch update.Message.Command() {
			case "start":
				reply = "Hello!"
			case "help":
				reply = `News are updated every 10 seconds
Possible commands: 
  /add_rss url
    Example: /add_rss http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml
  /add_web url item-css-selector link-css-selector title-css-selector
    Example /add_web https://lenta.ru/parts/news ".item.news" ".titles a" ".titles a"
    item-css-selector is for selecting news item on a page
    link-css-selector is for selecting <a> within news item
    title-css-selector is for selecting html tag (inside item) which has title string as inner text 
  /fetchnews [title]
  /stop
`
			case "settings":
				reply = "No settings"
			case "add_rss":
        url := update.Message.CommandArguments()
        if len(url) == 0 {
          reply = "no url provided for /add_rss"
        } else {
          debug("RSS url: " + url)
          rs, err := rssSrc.NewRssNewsSource("http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
          if err != nil {
            reply = "Failed to add RSS source from " + url
            debug(reply)
          } else {
            newsSrv.AddSource(rs)
            reply = "Added RSS source"
          }
        }
			case "add_web":
        args, err := shellwords.Parse(update.Message.CommandArguments())
        if len(args) != 4 || err != nil {
          reply = "failed to parse 4 arguments from /add_web"
        } else {
          debug("Web url: " + args[0])
          hs, err := httpSrc.NewHttpNewsSource(args[0])
          if err != nil {
            reply = "Failed to add web source from " + args[0]
            debug(reply)
          } else {
            hs.SetParser(map[string]string{
              "item": args[1],
              "link": args[2],
              "title": args[3],
            })
            newsSrv.AddSource(hs)
            reply = "Added Web source"
          }
        }
			case "fetchnews":
        title := update.Message.CommandArguments()
        debug("Title argument: " + title)
        articles, _ := newsSrv.Repo.GetArticlesBy(title)
        debug(fmt.Sprintf("About to send %d articles", len(articles)))
        for _, a := range articles {
          msg := tgbotapi.NewMessage(update.Message.Chat.ID, a.Title + " " + a.Link)
          bot.Send(msg)
        }
        reply = fmt.Sprintf("no more articles, %d total", len(articles))
			case "stop":
				newsSrv.Stop()
				return nil
			}
		}
    debug("REPLY: " + reply)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}

	return nil
}
