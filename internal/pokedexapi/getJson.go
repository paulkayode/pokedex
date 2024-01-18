package pokedexapi

import (
	"io"
	"net/http"
	"github.com/segunkayode1/pokedex/internal/cache"
	"time"
	"fmt"
)
const interval = 5 * time.Second

var m_cache cache.Cache = cache.NewCache(interval)

func GetJson(url string) ([]byte, error){
	if val, ok := m_cache.Get(url); ok{
		return val, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	val, err:= io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 399 {
		return []byte{}, fmt.Errorf("network-issue, status code: %v", resp.Status)
	}
	if err != nil {
		return []byte{}, err
	}
	m_cache.Put(url, val)
	return val, nil
	
}