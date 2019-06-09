module github.com/danielbintar/angel/server/users

go 1.12

require (
	github.com/danielbintar/angel/server-library v0.0.0-20190529145453-576c2ee29bbf
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.8
	github.com/julienschmidt/httprouter v1.2.0
	github.com/stretchr/testify v1.3.0
	github.com/subosito/gotenv v1.1.1
	golang.org/x/crypto v0.0.0-20190404164418-38d8ce5564a5
	gopkg.in/validator.v2 v2.0.0-20180514200540-135c24b11c19
)

replace github.com/danielbintar/angel/server-library => ../../server-library
