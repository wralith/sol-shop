module github.com/wralith/sol-shop/user-service

go 1.20

require (
	github.com/wralith/sol-shop/pb v0.0.0
	golang.org/x/crypto v0.10.0
	google.golang.org/grpc v1.55.0
	gorm.io/driver/sqlite v1.5.1
	gorm.io/gorm v1.25.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/wralith/sol-shop/pb v0.0.0 => ../pb
