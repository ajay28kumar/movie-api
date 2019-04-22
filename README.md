# movie-api
a sample of movie data using goLang

### installation
1. Install [dep](https://github.com/golang/dep)
2. Run `dep ensure` in project root directory to install all dependencies
3. create `.env` folder to set environment
 
 .env
```$xslt
        HOST="localhost"
        PORT=5432
        USER="postgres"
        PASSWORD="******"
        DBNAME="postgres"
```
4. Start server `go run main.go`

