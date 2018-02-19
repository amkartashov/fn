package main

import (
	"fmt"
	//news "github.com/gorilych/fn/news"
	//aRepo "github.com/gorilych/fn/news/articlerepo/sqlite"
	httpSrc "github.com/gorilych/fn/news/source/http"
	rssSrc "github.com/gorilych/fn/news/source/rss"
)

func main() {
	//var a = news.Article{Title: "Lorem Ipsum", Link: "https:/example.com" }
	//repo, err := aRepo.NewSqliteArticleRepo("./articles.db")
	//if err != nil { fmt.Println(err) }
	//fmt.Println(a)
	//if err == nil {
	//  repo.StoreArticle(a)
	//}
	//as, err := repo.GetArticlesBy("")
	//if err != nil { fmt.Println(err) }
	//fmt.Println(as)
	rs, err := rssSrc.NewRssNewsSource("http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
	if err != nil {
		fmt.Println(err)
	}
	as, _ := rs.FetchAllArticles()
	fmt.Println(as)
	hs, err := httpSrc.NewHttpNewsSource("https://lenta.ru/parts/news")
	if err != nil {
		fmt.Println(err)
	}
	hs.SetParser(map[string]string{
		"item":  ".item.news",
		"link":  ".titles a",
		"title": ".titles a",
	})
	as, err = hs.FetchAllArticles()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(as)
}
