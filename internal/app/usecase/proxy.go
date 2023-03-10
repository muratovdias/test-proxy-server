package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/muratovdias/test-proxy-server/internal/app/cache"
	"github.com/muratovdias/test-proxy-server/internal/app/entities"
)

type ProxyUsecase struct {
	cache *cache.Cache
}

func NewProxyUsecase(cache *cache.Cache) *ProxyUsecase {
	return &ProxyUsecase{
		cache: cache,
	}
}

func (p *ProxyUsecase) ProxyRequest(request entities.ProxyRequest) (entities.ProxyResponse, error) {
	if request.Method == "POST" {
		return entities.ProxyResponse{}, fmt.Errorf("only GET method is allowed")
	}
	_, err := url.Parse(request.URL)
	if err != nil {
		return entities.ProxyResponse{}, fmt.Errorf("usecase: ProxyRequest: Parse url: %w", err)
	}
	cacheKey, err := makeKeyForCache(request)
	if err != nil {
		return entities.ProxyResponse{}, fmt.Errorf("usecase: ProxyRequest: makeKeyForCache: %w", err)
	}
	response, ok := p.cache.Get(cacheKey) // check cache, if request have already been, just return it from cahce
	if ok {
		log.Println("response from cache")
		return response, nil
	}
	newRequest, err := http.NewRequest(request.Method, request.URL, nil) // create new request
	if err != nil {
		return entities.ProxyResponse{}, fmt.Errorf("usecase: ProxyRequest: NewRequest: %w", err)
	}
	for v, k := range request.Headers { // set headers for new request
		newRequest.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(newRequest) // get response
	if err != nil {
		return entities.ProxyResponse{}, fmt.Errorf("usecase: ProxyRequest: client.Do: %w", err)
	}

	p.cache.Set(cacheKey, makeProxyResponse(resp))
	return makeProxyResponse(resp), nil
}

// makeKeyForCache modifys request to string (key)
func makeKeyForCache(request entities.ProxyRequest) (string, error) {
	headersBytes, err := json.Marshal(request.Headers)
	if err != nil {
		return "", fmt.Errorf("usecase: makeKeyForCache: %w", err)
	}
	headerString := string(headersBytes)
	return fmt.Sprintf("%s:%s:%s", request.Method, request.URL, headerString), nil
}

func makeProxyResponse(response *http.Response) entities.ProxyResponse {
	var proxyResponse entities.ProxyResponse
	proxyResponse.Headers = make(map[string][]string)
	proxyResponse.ID = uuid.New().String()
	proxyResponse.Status = response.Status
	proxyResponse.Length = int(response.ContentLength)
	for k, v := range response.Header {
		proxyResponse.Headers[k] = v
	}
	return proxyResponse
}
