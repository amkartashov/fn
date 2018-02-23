package main

import (
	"fmt"
	ui "github.com/gorilych/fn/news/ui/telegram"
	// httpSrc "github.com/gorilych/fn/news/source/http"
	// rssSrc "github.com/gorilych/fn/news/source/rss"
	aRepo "github.com/gorilych/fn/news/articlerepo/sqlite"
	//"time"
	"os"
)

func main() {
	//var a = news.Article{Title: "Lorem Ipsum", Link: "https:/example.com" }
	token := os.Getenv("TGBOTTOKEN")
	repo, err := aRepo.NewSqliteArticleRepo("./articles.db")
	if err != nil {
		fmt.Println(err)
	}

	ui.Run(repo, token)
	// rs, err := rssSrc.NewRssNewsSource("http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
	// if err != nil {
	//   fmt.Println(err)
	// }
	// as, _ := rs.FetchAllArticles()
	// fmt.Println("============ Articles in RSS")
	// fmt.Println(len(as))
	// hs, err := httpSrc.NewHttpNewsSource("https://lenta.ru/parts/news")
	// if err != nil {
	//   fmt.Println(err)
	// }
	// hs.SetParser(map[string]string{
	//   "item":  ".item.news",
	//   "link":  ".titles a",
	//   "title": ".titles a",
	// })
	// as, err = hs.FetchAllArticles()
	// if err != nil {
	//   fmt.Println(err)
	// }
	// fmt.Println("============ Articles in HTML")
	// fmt.Println(len(as))
	// newsSrv := news.NewsService(repo, 2*time.Second)
	// fmt.Println("============ Started service")
	// newsSrv.AddSource(rs)
	// fmt.Println("============ Added RSS, now wait")
	// time.Sleep(3 * time.Second)
	// as, _ = repo.GetArticlesBy("")
	// fmt.Println(len(as))
	// newsSrv.AddSource(hs)
	// fmt.Println("============ Added HTTP, now wait")
	// time.Sleep(3 * time.Second)
	// as, _ = repo.GetArticlesBy("")
	// fmt.Println(len(as))
}
