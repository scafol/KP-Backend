# Discussion API

## golang version 
- 1.5

## library requirements
- [echo](https://echo.labstack.com/)
- [gorm](https://gorm.io/)
- [gorm postgres driver](http://gorm.io/driver/postgres)
- [godotenv](https://github.com/joho/godotenv)
- [logrus](https://github.com/sirupsen/logrus)

## Environment variable declaration

1. Define .env file
    ```shell
        $~ mv .env.example .env
    ```
2. Define CONN_STRING in .env file, for example :    
    ```shell
        CONN_STRING="user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    ```
3. define LISTEN_PORT in .env file, for example : 
    ```shell
        LISTEN_PORT=":8080"

## Installation
### Using Docker    
 - running docker-compose
    ```shell
        $~ docker-compose up --build -d
    ```
### Without Docker
 - Run directly in shell
    ```shell
        $~ go run ./main.go
    ```    
