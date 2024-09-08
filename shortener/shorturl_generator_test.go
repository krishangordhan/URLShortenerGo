package shortener_test

import (
	"testing"

	"github.com/krishangordhan/go-url-shortener/shortener"
)

const UserID = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink1 := shortener.GenerateShortLink(initialLink1, UserID)

	initialLink2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink2 := shortener.GenerateShortLink(initialLink2, UserID)

	initialLink3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink3 := shortener.GenerateShortLink(initialLink3, UserID)

	if shortLink1 != "jTa4L57P" {
		t.Errorf("ShortLinkGenerator() failed for initialLink_1; want: jTa4L57P, got: %s", shortLink1)
	}

	if shortLink2 != "d66yfx7N" {
		t.Errorf("ShortLinkGenerator() failed for initialLink_2; want: d66yfx7N, got: %s", shortLink2)
	}

	if shortLink3 != "dhZTayYQ" {
		t.Errorf("ShortLinkGenerator() failed for initialLink_3; want: dhZTayYQ, got: %s", shortLink3)
	}
}
