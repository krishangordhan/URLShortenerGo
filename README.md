# URLShortenerGo

Has three stages to this:

 - 1. Shorten URL using sha2560f and Base58 encoding and grabbing the first 8 characters.
 - 2. Store the shortened URL in a map with the shortened URL as the key and the original URL as the value in Redis server.
 - 3. Redirect the shortened URL to the original URL.