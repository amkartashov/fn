package http

import (
  "errors"
  "strings"
  "regexp"
  news "github.com/gorilych/fn/news"
  gq "github.com/PuerkitoBio/goquery"
)

type httpNewsSource struct {
  Address string
  selectors map[string]string
}

func NewHttpNewsSource(url string) (httpNewsSource, error) {
  src := httpNewsSource{Address: url, selectors: make(map[string]string)}
  return src, nil
}

func (s httpNewsSource) GetParserArgs() (map[string]string, error) {
  return map[string]string{
    "item": "css selector for news block with one article",
    "link": "css selector within a news block for <a> tag with link to article",
    "title": "css selector within a news block for block which internal html is a title",
  }, nil
}

func (s httpNewsSource) SetParser(args map[string]string) error {
  argsdescr, err := s.GetParserArgs()
  if err != nil { return err }
  for argname, descr := range argsdescr {
    val, ok := args[argname]
    if !ok {
      return errors.New("Parser argument " + argname + " (" + descr + ") is not passed to SetParser()")
    }
    s.selectors[argname] = val
  }
  return nil
}

func (s httpNewsSource) FetchAllArticles() ([]news.Article, error) {
  var as []news.Article
  doc, err := gq.NewDocument(s.Address)
  if err != nil { return as, err }
  doc.Find(s.selectors["item"]).Each(func(index int, item *gq.Selection) {
    link, _ := item.Find(s.selectors["link"]).Attr("href")
    link = href2url(s.Address, link)
    title := item.Find(s.selectors["title"]).Text()
    as = append(as, news.Article{Title: title, Link: link})
  })
  return as, nil
}

// Given base URL and href value, produce full URL for a link
func href2url(docUrl string, href string) string {
  if (strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://")) {
    return href
  }
  if strings.HasPrefix(href, "/") {
    r := regexp.MustCompile(`https?://[^/]+`)
    baseUrl := r.FindAllString(docUrl, -1)[0]
    return baseUrl + href
  }
  return docUrl + "/" + href
}








