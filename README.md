# golang-Api-sample

# Create a GOPATH
```export GOPATH=$USERPROFILE/go/ ```

# Create a project folder

```mkdir -p $USERPROFILE/go/src/github.com/{your project name} ```

# Create a file named api server

``` cd {your project name} ```

``` touch api.go ```

copy all the codes from totalmarger.go to the api.go

# Install dependencies 

```go get -u github.com/go-sql-driver/mysql github.com/gorilla/mux```


# Create a MYSQL database named test

Create a MySQL database test and replace the ```database user``` and ```password``` with your Mysql user and password in api.go file

# Now build 

```` go build ```

it will create a {your project name}.exe file

# Now run that file

``` ./{your project name} ```

# OR run directly

``` go run api.go ```







