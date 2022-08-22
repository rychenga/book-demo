# book-demo
borrow books system demo
::功能項目
1. gin 架設 Web Service
2. orm 設計 「物件關係的映射」對資料庫執行 新增、修改、刪除及查詢 功能
3. Web Service 採用 RESTful 作為設計，完成一項功能的 CRUD(Create、Read、Update、Delete)。
4. 資料庫採用SQLite

::impoart
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u github.com/mattn/go-oci8

::指定功能
:GET:    Read
curl localhost:8080/GET
http://localhost:8080/GET?name=A_BOOK

::POST:   Create
curl localhost:8080/POST
http://localhost:8080/POST?name=A_BOOK
http://localhost:8080/POST?name=B_BOOK
http://localhost:8080/POST?name=C_BOOK
http://localhost:8080/POST?name=D_BOOK
http://localhost:8080/POST?name=E_BOOK

::PUT:    Update
curl localhost:8080/PUT
http://localhost:8080/PUT?keyname=A_BOOK&borrower=RYCHENGA
http://localhost:8080/PUT?keyname=A_BOOK&borrower=%20

::DELETE: Delete
curl localhost:8080/DELETE
http://localhost:8080/DELETE?name=E_BOOK

