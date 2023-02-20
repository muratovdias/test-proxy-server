package usecase

import (
	"github.com/muratovdias/test-proxy-server/internal/app/cache"
	"github.com/muratovdias/test-proxy-server/internal/app/entities"
)

type Usecase struct {
	Proxy
}

type Proxy interface {
	ProxyRequest(entities.ProxyRequest) (entities.ProxyResponse, error)
}

func NewUsecase(cache *cache.Cache) *Usecase {
	return &Usecase{
		Proxy: NewProxyUsecase(cache),
	}
}
