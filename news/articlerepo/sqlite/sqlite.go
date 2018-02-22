package sqlite

import (
	sql "database/sql"
	news "github.com/gorilych/fn/news"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteArticleRepo struct {
	FilePath string
	DB       *sql.DB
	saveStmt *sql.Stmt
	getStmt  *sql.Stmt
}

func NewSqliteArticleRepo(FilePath string) (sqliteArticleRepo, error) {
	repo := sqliteArticleRepo{FilePath: FilePath}
	db, err := sql.Open("sqlite3", FilePath)
	if err != nil {
		return repo, err
	}
	repo.DB = db
	statement, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS article (
      id INTEGER PRIMARY KEY,
      title TEXT,
      link TEXT,
      UNIQUE (title, link)
    )`)
	if err != nil {
		return repo, err
	}
	statement.Exec()
	repo.saveStmt, err = db.Prepare("INSERT OR IGNORE INTO article(title, link) VALUES (?, ?)")
	if err != nil {
		return repo, err
	}
	repo.getStmt, err = db.Prepare("SELECT title, link FROM article WHERE title LIKE ?")
	if err != nil {
		return repo, err
	}
	return repo, nil
}

// sqliteArticleRepo needs to implement interface news.ArticleRepo

func (r sqliteArticleRepo) StoreArticle(a news.Article) error {
	_, err := r.saveStmt.Exec(a.Title, a.Link)
	return err
}

func (r sqliteArticleRepo) GetArticlesBy(s string) ([]news.Article, error) {
	var as []news.Article
	rows, err := r.getStmt.Query("%" + s + "%")
	if err != nil {
		return as, err
	}
	for rows.Next() {
		var a news.Article
		err = rows.Scan(&a.Title, &a.Link)
		if err != nil {
			return as, err
		}
		as = append(as, a)
	}
	return as, nil
}
