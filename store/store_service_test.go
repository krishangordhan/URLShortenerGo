package store_test

import (
	"testing"

	"github.com/krishangordhan/go-url-shortener/store"
)

var testStoreService *store.StorageService

func init() {
	testStoreService = store.InitializeStore()
}

func TestStoreInit(t *testing.T) {
	if testStoreService.RedisClient == nil {
		t.Errorf("InitalizeStore() should return a valid redis client; want: not nil, got: nil")
	}
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://github.com/krishangordhan/go-url-shortener"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	store.SaveURLMapping(shortUrl, initialLink, userUUID)

	retrievedUrl := store.RetrieveURL(shortUrl)

	if retrievedUrl != initialLink {
		t.Errorf("RetrieveUrl() should return the original url; want: %s, got: %s", initialLink, retrievedUrl)
	}
}
