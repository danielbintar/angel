module github.com/danielbintar/angel/server/users

go 1.12

require (
	github.com/DataDog/zstd v1.3.8 // indirect
	github.com/danielbintar/angel/server-library v0.0.0-20190617035758-93f32450d08e
	github.com/go-sql-driver/mysql v1.4.1
	github.com/julienschmidt/httprouter v1.2.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/validator.v2 v2.0.0-20180514200540-135c24b11c19
)

replace github.com/danielbintar/angel/server-library => ../../server-library
