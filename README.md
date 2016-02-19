# go-nowplaying
Tweet iTunes's nowplaying song for **Windows and OSX**.


## Install
```
go get github.com/yaegaki/go-nowplaying
```

## Usage
### Move go-nowplaying directory
```sh
cd $GOPATH/src/github.com/yaegaki/go-nowplaying
```
### Edit Consumer Key and etc of main.go
```go
// set your values
const ConsumerKey = "Your Consumer Key (API Key)"
const ConsumerSecret = "Your Consumer Secret (API Secret)"
const AccessToken = "Your Access Token"
const AccessTokenSecret = "Your Access Token Secret"
```
  
### Run
```sh
go run main.go
```

## License
### MIT
