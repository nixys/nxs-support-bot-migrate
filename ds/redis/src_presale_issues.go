package redis

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/nixys/nxs-support-bot-migrate/misc"

	"github.com/go-redis/redis"
)

const srcPresaleIssuesKey = "nxs-chat-srv:nixys:presale"

func (r *Redis) SrcPresalesGet() (map[int64]int64, error) {

	iss := make(map[int64]int64)

	irs := r.client.HGetAll(srcPresaleIssuesKey)
	if irs.Err() != nil {
		if irs.Err() == redis.Nil {
			// Empty keys
			return nil, fmt.Errorf("redis src presale issues get: %w", misc.ErrNotFound)
		}
		return nil, fmt.Errorf("redis src presale issues get: %w", irs.Err())
	}

	for k, val := range irs.Val() {

		key, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("redis src presale issues get: %w", err)
		}

		var v int64
		if err := json.Unmarshal([]byte(val), &v); err != nil {
			return nil, fmt.Errorf("redis src presale issues get: %w", err)
		}
		iss[key] = v
	}

	return iss, nil
}
