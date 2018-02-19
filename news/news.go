// Package news provides DTO for business entities
// and high level interfaces

package news

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
