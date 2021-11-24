package cacheinmem

import (
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
)

const (
	TTL = 60
)

var (
	Client cacheClientInterface = &cacheClient{}
)

type cacheClientInterface interface {
	setClient(incache *ttlcache.Cache, CacheResp []*InMemoryPerformances)
	GetINuser(hash string) bool
	SetINuser(hash string, query queries.EsQuery)
	Erase()
	SettOutUser(hash string, performan []performance.Performance)
	GetOutUser(hash string) []performance.Performance
}
type cacheClient struct {
	CacheIn   *ttlcache.Cache
	CacheResp []*InMemoryPerformances
}

type InMemoryPerformances struct {
	performances []performance.Performance
	hashP        string
}

func Init() {
	in := ttlcache.NewCache()

	in.SetTTL(time.Duration(TTL * time.Second))

	Client.setClient(in, make([]*InMemoryPerformances, 0))

}

func (c *cacheClient) setClient(incache *ttlcache.Cache, CacheResp []*InMemoryPerformances) {
	c.CacheIn = incache
	c.CacheResp = CacheResp
}

func (ci *cacheClient) GetINuser(hash string) bool {

	if _, err := ci.CacheIn.Get(hash); err != ttlcache.ErrNotFound {
		return true
	}
	return false
}

func (ci *cacheClient) SetINuser(hash string, query queries.EsQuery) {
	ci.CacheIn.Set(hash, query)

}

func (ci *cacheClient) GetOutUser(hash string) []performance.Performance {
	for _, v := range ci.CacheResp {
		if v.hashP == hash {
			return v.performances
		}
	}
	return nil
}

func (ci *cacheClient) SettOutUser(hash string, performan []performance.Performance) {

	container := &InMemoryPerformances{
		performances: performan,
		hashP:        hash,
	}

	ci.CacheResp = append(ci.CacheResp, container)

}

func (ci *cacheClient) Erase() {
	ci.CacheIn.Purge()

	ci.CacheResp = nil

}
