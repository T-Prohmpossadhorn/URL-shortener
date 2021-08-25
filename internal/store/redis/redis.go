package redis

import (
	"fmt"
	"time"

	"url-shortener/internal/shortener"
	"url-shortener/internal/store"

	redisClient "github.com/gomodule/redigo/redis"
)

type redis struct {
	pool     *redisClient.Pool
	useruuid string
}

func New(host, port, useruuid string) (store.Service, error) {
	pool := &redisClient.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redisClient.Conn, error) {
			return redisClient.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}

	return &redis{pool, useruuid}, nil
}

func (r *redis) Save(url string, expires time.Time) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	shortlink, err := shortener.GenerateShortLink(url, r.useruuid)
	if err != nil {
		return "", err
	}

	_, err = conn.Do("HMSET", shortlink, "link", shortlink, "url", url, "expires", expires.Format("Mon, 02 Jan 2006 15:04:05 MST"), "visits", 0)
	if err != nil {
		return "", err
	}

	return shortlink, nil
}

func (r *redis) Load(shortlink string) (string, error) {
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

	expire, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", shortLink.Expires)
	if err != nil {
		return "", err
	} else if expire.Before(time.Now()) {
		return "", fmt.Errorf("url expired")
	}

	conn.Do("HINCRBY", shortlink, "visits", 1)

	return shortLink.URL, nil
}

func (r *redis) LoadInfo(shortlink string) (*store.Item, error) {
	conn := r.pool.Get()
	defer conn.Close()

	values, err := redisClient.Values(conn.Do("HGETALL", shortlink))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, fmt.Errorf("url not found")
	}

	var shortLink store.Item
	err = redisClient.ScanStruct(values, &shortLink)
	if err != nil {
		return nil, err
	}

	return &shortLink, nil
}

func (r *redis) Close() error {
	return r.pool.Close()
}
