# lectio-api


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
        "userId": 4,
        "bookName": "オーバーロード",
        "author": "先生",
        "publisher": "出版社",
        "price": 850,
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
        "bookName": "オーバーロード2",
        "author": "先生2",
        "publisher": "出版社2",
        "price": 2850,
        "bookStatus": "favorite2"
    }' \
'localhost:8000/v1/books/14'
```

* Delete
```sh
curl -X DELETE localhost:8000/v1/books/14
```