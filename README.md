# GOLANG CRUD - WEBSITE
 Pemrograman CRUD Dengan Tampilan Website
 
## Tech Stack
1. MySQL
2. Golang
3. External packages :
* XAMPP
* Bootsrap v5.3

## Setting Port :
```
file : main.go
	http.ListenAndServe(":['change youre port here']", nil)
```
## Setting Database :
```
file : config/database.go
	db, err := sql.Open("mysql", "root:@/['change youre database here']?parseTime=true")
```

## Prerequisites
Enter the following command to install sql driver for mysql, gin and gorm in the project.
```
go get -u github.com/go-sql-driver/mysql
```
or you can go here to see how to use it :
```
https://github.com/go-sql-driver/mysql
```

## Running
```
go run main.go
```

## CRUD Endpoints
1. PRODUCTS
* /products/add - GET & POST
* /products/detail?id - GET
* /products/edit?id - GET & POST
* /products/delete?id - POST

3. CATEGORIES
* /categories/add - GET & POST
* /categories/edit - GET & POST
* /categories/delete - POST
