package rss

import ( 
  news "github.com/gorilych/fn/news"
  gofeed "github.com/mmcdole/gofeed"
)

type rssNewsSource struct {
  Address string
  feed *gofeed.Feed
}

func NewRssNewsSource(url string) (rssNewsSource, error) {
  src := rssNewsSource{Address: url}
  fp := gofeed.NewParser()
  feed, err := fp.ParseURL(url)
  if err != nil { return src, err }
  src.feed = feed
  return src, nil
}

func (s rssNewsSource) GetParserArgs() (map[string]string, error) {
  return map[string]string{}, nil
}

func (s rssNewsSource) SetParser(args map[string]string) error {
  return nil
}

func (s rssNewsSource) FetchAllArticles() ([]news.Article, error) {
  var as []news.Article
  for _, i := range s.feed.Items {
    as = append(as, news.Article{i.Title, i.Link})
  }
  return as, nil
}


