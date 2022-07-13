package auth

type SimpleNonceCache struct {
	Data map[string][]byte
}

func CreateSimpleNonceCache() *SimpleNonceCache {

	return &SimpleNonceCache{
		Data: make(map[string][]byte),
	}

}

func (cache *SimpleNonceCache) Get(key string) ([]byte, error) {
	return cache.Data[key], nil
}

func (cache *SimpleNonceCache) Put(key string, data []byte) error {
	cache.Data[key] = data
	return nil
}
