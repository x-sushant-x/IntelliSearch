### Simple Web Crawler Written in Golang

It not a perfect implementation. I made it because I was bored of making APIs and wanted to learn something new. 

### Steps To Run:

1. Create following kafka topics: - `crawl_urls`, `crawled_pages`
2. Inside `crawl_manager` and `crawler` folder run `make run`
3. Hit this cURL

```
curl --location 'localhost:8080/crawl' \
--header 'Content-Type: application/json' \
--data '{
    "urls" : ["URL HERE", "URL HERE"]
}'
```

### Want to change topic name?
1. Change them in `crawl_manager/cmd/main.go` and `crawler/cmd/main.go`
2. Or please user [Viper](https://github.com/spf13/viper) and open a raise a Pull Request 🙃.

### Want to improve this?
Always welcome.