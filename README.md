# lectio-api

### Goaの利用
```sh
cd design;goa gen github.com/Shimizu1111/lectio/design;cd ../;go build ./design/cmd/lectio && go build ./design/cmd/lectio-cli;./lectio
```

### APIの例
* Read
```sh
curl -X GET \
    -H "Content-Type: application/json" \
    -d \
    '{
        "bookId": 1
    }' \
'localhost:8000/v1/books/1'
```

* ReadAll
```sh
curl -X GET localhost:8000/v1/books
```

* Create
```sh
curl -X POST \
    -H "Content-Type: application/json" \
    -d \
    '{
        "userId": 1,
        "bookName": "坊っちゃん",
        "author": "夏目漱石",
        "publisher": "角川文庫",
        "price": 500,
        "bookStatus": "favorite"
    }' \
'localhost:8000/v1/books'
```

* Update
```sh
curl -X PUT \
    -H "Content-Type: application/json" \
    -d \
    '{
        "bookName": "吾輩は猫である",
        "author": "夏目漱石",
        "publisher": "新潮文庫",
        "price": 700,
        "bookStatus": "favorite"
    }' \
'localhost:8000/v1/books/14'
```

* Delete
```sh
curl -X DELETE localhost:8000/v1/books/14
```