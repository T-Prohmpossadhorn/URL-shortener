package redis

import (
	"fmt"
	"time"

	"url-shortener/config"
	"url-shortener/internal/shortener"
	"url-shortener/internal/store"

	redisClient "github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool     *redisClient.Pool
	useruuid string
}

func New(config config.Config) (store.Service, error) {
	pool := &redisClient.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redisClient.Conn, error) {
			return redisClient.Dial("tcp", fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port))
		},
	}

	return &Redis{pool, config.JwtAccount.UUID}, nil
}

func (r *Redis) Save(url string, expires string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	shortlink, err := shortener.GenerateShortLink(url, r.useruuid)
	if err != nil {
		return "", err
	}

	_, err = conn.Do("HMSET", shortlink, "link", shortlink, "url", url, "expires", expires, "visits", 0)
	if err != nil {
		return "", err
	}

	return shortlink, nil
}

func (r *Redis) Load(shortlink string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	values, err := redisClient.Values(conn.Do("HGETALL", shortlink))
	if err != nil {
		return "nil", err
	} else if len(values) == 0 {
		return "", fmt.Errorf("url not found")
	}

	var shortLink store.Item
	err = redisClient.ScanStruct(values, &shortLink)
	if err != nil {
		return "", err
	}

	if len(shortLink.URL) == 0 {
		return "", fmt.Errorf("url not found")
	}

	if shortLink.Expires != "" {
		expire, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", shortLink.Expires)
		if err != nil {
			return "", err
		} else if expire.Before(time.Now()) {
			return "", fmt.Errorf("url expired")
		}
	}

	conn.Do("HINCRBY", shortlink, "visits", 1)

	return shortLink.URL, nil
}

func (r *Redis) LoadInfo(shortlink string) (*store.Item, error) {
	conn := r.pool.Get()
	defer conn.Close()

	values, err := redisClient.Values(conn.Do("HGETALL", shortlink))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, fmt.Errorf("url not found")
	}

	var data store.Item
	err = redisClient.ScanStruct(values, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Redis) Delete(shortlink string) error {
	conn := r.pool.Get()
	defer conn.Close()

	values, err := conn.Do("EXISTS", shortlink)
	if err != nil {
		return err
	} else if values == 0 {
		return fmt.Errorf("cannot find short link")
	}
	// expired from now on
	_, err = conn.Do("HSET", shortlink, "expires", time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
	return err
}

func (r *Redis) Close() error {
	return r.pool.Close()
}
