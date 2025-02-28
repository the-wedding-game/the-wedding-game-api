package utils

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestIsURLStrictWithValidUrls(t *testing.T) {
	assert.Equal(t, IsURLStrict("http://www.google.com"), true)
	assert.Equal(t, IsURLStrict("https://www.google.com"), true)
	assert.Equal(t, IsURLStrict("http://google.com"), true)
	assert.Equal(t, IsURLStrict("https://google.com"), true)
	assert.Equal(t, IsURLStrict("http://www.google.com/"), true)
	assert.Equal(t, IsURLStrict("https://www.google.com/"), true)
	assert.Equal(t, IsURLStrict("http://google.com/"), true)
	assert.Equal(t, IsURLStrict("https://google.com/"), true)
	assert.Equal(t, IsURLStrict("http://www.google.com/path"), true)
	assert.Equal(t, IsURLStrict("https://www.google.com/path"), true)
	assert.Equal(t, IsURLStrict("http://google.com/path"), true)
	assert.Equal(t, IsURLStrict("https://google.com/path"), true)
	assert.Equal(t, IsURLStrict("https://google.com/path/image.jpg"), true)
	assert.Equal(t, IsURLStrict("https://google.net/path/image.jpg"), true)
	assert.Equal(t, IsURLStrict("https://subdomain.google.net/path/image.jpg"), true)
	assert.Equal(t, IsURLStrict("http://localhost:9445/the-wedding-game/testing-folder/d3d2df12-f5dc-11ef-a7f9-7acdc90419c1k.png"), true)
}

func TestIsURLStrictWithInvalidUrls(t *testing.T) {
	assert.Equal(t, IsURLStrict("www.google.com"), false)
	assert.Equal(t, IsURLStrict("google.com"), false)
	assert.Equal(t, IsURLStrict("www.google.com/"), false)
	assert.Equal(t, IsURLStrict("google.com/"), false)
	assert.Equal(t, IsURLStrict("www.google.com/path"), false)
	assert.Equal(t, IsURLStrict("google.com/path"), false)
	assert.Equal(t, IsURLStrict("google.com/path/image.jpg"), false)
	assert.Equal(t, IsURLStrict("google.net/path/image.jpg"), false)
	assert.Equal(t, IsURLStrict("subdomain.google.net/path/image.jpg"), false)
	assert.Equal(t, IsURLStrict("http://"), false)
	assert.Equal(t, IsURLStrict("https://"), false)
	assert.Equal(t, IsURLStrict("https://invalid"), false)
	assert.Equal(t, IsURLStrict("www.google.com"), false)
}
