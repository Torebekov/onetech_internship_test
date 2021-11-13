package acmp_concurrent

import (
	"github.com/Torebekov/onetech_internship_test/acmp"
	"sync"
)

var (
	cache = sync.Map{}
	wg    sync.WaitGroup
)

func Difficulties(urls []string) map[string]float64 {
	mapDiff := make(map[string]float64)
	for _, url := range urls {
		wg.Add(1)
		go getDifficulty(url)
	}
	wg.Wait()
	cache.Range(func(key interface{}, value interface{}) bool {
		mapDiff[key.(string)] = value.(float64)
		return true
	})
	return mapDiff
}
func getDifficulty(url string) {
	defer wg.Done()
	cache.Store(url, acmp.Difficulty(url))
}
