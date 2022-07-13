package ext

type CacheService interface {
	Get(key string) ([]byte, error)
	Put(key string, data []byte) error
}
