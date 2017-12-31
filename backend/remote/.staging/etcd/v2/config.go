package etcd2cache

import (
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
	"github.com/sniperkit/configer"
	"github.com/sniperkit/httpcache/pkg"
	"github.com/sniperkit/vipertags"
)

type etcd2cacheConfig struct {
	Provider       string        `json:"provider" yaml:"provider" config:"cache.http.provider"`
	MaxConnections int           `json:"max_connections" yaml:"max_connections" config:"cache.http.max_connections" default:"0"`
	BucketName     string        `json:"bucket_name" yaml:"bucket_name" config:"cache.http.provider"`
	StoragePath    string        `json:"storage_path" yaml:"storage_path" config:"cache.http.storage_path"`
	ReadOnly       bool          `json:"read_only" yaml:"read_only" config:"cache.http.read_only"`
	StrictMode     bool          `json:"strict_mode" yaml:"strict_mode" config:"cache.http.strict_mode"`
	NoSync         bool          `json:"no_sync" yaml:"no_sync" config:"cache.http.no_sync"`
	NoFreelistSync bool          `json:"no_freelist_sync" yaml:"no_freelist_sync" config:"cache.http.no_freelist_sync"`
	NoGrowSync     bool          `json:"no_grow_sync" yaml:"no_grow_sync" config:"cache.http.no_grow_sync"`
	MaxBatchSize   bool          `json:"max_batch_size" yaml:"max_batch_size" config:"cache.http.max_batch_size"`
	MaxBatchDelay  bool          `json:"max_batch_delay" yaml:"max_batch_delay" config:"cache.http.max_batch_delay"`
	AllocSize      bool          `json:"alloc_size" yaml:"alloc_size" config:"cache.http.provialloc_sizeder"`
	done           chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	PluginConfig = &etcd2cacheConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (etcd2cacheConfig) ConfigName() string {
	return "Etcd2"
}

// SetDefaults ...
func (a *etcd2cacheConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *etcd2cacheConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.Provider == "" {
		a.Provider = a.ConfigName()
	}
	if a.MaxConnections == 0 {
		a.MaxConnections = httpcache.DefaultMaxConnections
	}
}

// Read several config files (yaml, json or env variables)
func (a *etcd2cacheConfig) Configor(files []string) {
	configor.Load(&PluginConfig, files...)
}

// Wait ...
func (c etcd2cacheConfig) Wait() {
	<-c.done
}

// String ...
func (c etcd2cacheConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c etcd2cacheConfig) Debug() {
	// log.Debug("Etcd2 PluginConfig = ", c)
}

func init() {
	config.Register(PluginConfig)
}
