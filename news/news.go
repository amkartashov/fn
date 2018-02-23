// Package news provides DTO for business entities
// and high level interfaces

package news

import (
	"sync"
	"time"
)

type Article struct {
	Title string
	Link  string
}

type ArticleRepo interface {
	StoreArticle(a Article) error
	GetArticlesBy(s string) ([]Article, error)
}

type NewsSource interface {
	// return arguments for parser, argname -> argdescr
	GetParserArgs() (map[string]string, error)
	// set arguments for parser, argname -> argvalue
	SetParser(args map[string]string) error
	// return all articles from the source
	FetchAllArticles() ([]Article, error)
}

func CollectArticles(repo ArticleRepo, sources []NewsSource) {
	for _, s := range sources {
		articles, _ := s.FetchAllArticles()
		for _, a := range articles {
			repo.StoreArticle(a)
		}
	}
}

type newsService struct {
	Repo    ArticleRepo
	sources []NewsSource
	sMu     sync.RWMutex
	stopC   chan struct{}
}

func (srv *newsService) Collect() {
	srv.sMu.RLock()
	defer srv.sMu.RUnlock()
	CollectArticles(srv.Repo, srv.sources)
}

func (srv *newsService) AddSource(s NewsSource) {
	srv.sMu.Lock()
	defer srv.sMu.Unlock()
	srv.sources = append(srv.sources, s)
}

func NewsService(repo ArticleRepo, period time.Duration) *newsService {
	var srv newsService
	srv.Repo = repo
	srv.sources = make([]NewsSource, 0)
	srv.stopC = make(chan struct{})
	go func(srv *newsService) {
		srv.Collect()
		t := time.NewTicker(period)
		for {
			select {
			case <-srv.stopC:
				t.Stop()
				return
			case <-t.C:
				srv.Collect()
			}
		}
	}(&srv)
	return &srv
}

func (srv *newsService) Stop() {
	if srv.stopC != nil {
		srv.stopC <- struct{}{}
	}
	srv.stopC = nil
}
