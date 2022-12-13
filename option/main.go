package main

import "sync"

func main() {
	NewCache(SetShardCount(200))

}

type Cache struct {
	// hashFunc represents used hash func
	// HashFunc HashFunc
	// bucketCount represents the number of segments within a cache instance. value must be a power of two.
	BucketCount uint64
	// bucketMask is bitwise AND applied to the hashVal to find the segment id.
	bucketMask uint64
	// segment is shard
	// segments []*segment
	// segment lock
	locks []sync.RWMutex
	// close cache
	close chan struct{}
}

type Opt func(options *Cache)

func NewCache(opts ...Opt) (*Cache, error) {
	c := &Cache{
		close: make(chan struct{}),
	}

	for _, each := range opts {
		each(c)
	}
	return c, nil
}

func SetShardCount(count uint64) Opt {
	return func(opt *Cache) {
		opt.BucketCount = count
	}
}

// uber版本

type Options struct {
	bucketCount uint64
}

type Option interface {
	apply(options *Options)
}

type Bucket struct {
	count uint64
}

func (b Bucket) apply(opts *Options) {
	opts.bucketCount = b.count
}

func WithBucketCount(count uint64) Option {
	return Bucket{
		count: count,
	}
}

func NewOption(opts ...Option) (*Cache, error) {
	o := &Options{
		bucketCount: 20, // 默认
	}

	for _, each := range opts {
		each.apply(o)
	}
	return &Cache{
		BucketCount: o.bucketCount,
	}, nil
}
