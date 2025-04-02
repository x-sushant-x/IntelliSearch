package cache

type Cache interface {
	InsertBloomFilter(data string)
	CheckBloom(data string) bool
}
