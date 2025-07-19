package common

import "time"

type GetResponse struct {
	Value string `json:"value"`
}

type CacheSetRequest struct {
	Key   string        `json:"key"`
	Value interface{}   `json:"value"`
	TTL   time.Duration `json:"ttl,omitempty"`
}

type GetKeysResponse struct {
	Keys []string `json:"keys"`
}

type TtlResponse struct {
	Ttl int64 `json:"ttl"`
}
