package cache_service

import (
	"strconv"
	"strings"
	"test.go.dev/gin-blog/pkg/e"
)

type Article struct {
	ID int
	TagID int
	State int

	PageNum int
	PageSize int
}



func (a *Article) GetArticleKey() string {
	keys := [] string {
		e.CACHE_ARTICLE,
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	return strings.Join(keys, "_")
}


func (a *Article) GetArticleListKey() string {
	keys := [] string {
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	return strings.Join(keys, "_")
}