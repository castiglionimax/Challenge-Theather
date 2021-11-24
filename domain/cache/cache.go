package cache

import (
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
)

type CacheInUser struct {
	Query   queries.EsQuery
	IDCache int64
}

type CacheOUTUser struct {
	Performance []performance.Performance
	IDCache     int64
}
