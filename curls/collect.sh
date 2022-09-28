curl -X POST http://localhost:8080/collect \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id" : "771cd43f-011d-42fb-9aa6-2b45783e0836",
    "authentication_success" : true,
    "user_agent" : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
    "ip" : "186.192.87.8",
    "country" : "brazil",
    "region" : "south_america",
    "os" : "Mac OS X 10.15.7",
    "browser" : "Chrome",
    "device_type" : "Other 0.0.0"
}' -v