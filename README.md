# go-nowplaying
Tweet iTunes's nowplaying song for **Windows and OSX**.


## Install
```
go get github.com/yaegaki/go-nowplaying
```

## Usage
### Move go-nowplaying directory
for Windows
```cmd
cd %GOPATH%/src/github.com/yaegaki/go-nowplaying
```
for OSX
```sh
cd $GOPATH/src/github.com/yaegaki/go-nowplaying
```

### Set Env
for Windows
```cmd
set GN_CONSUMER_KEY=YOUR_CONSUMER_KEY
set GN_CONSUMER_SECRET=YOUR_CONSUMER_SECRET
set GN_ACCESS_TOKEN=YOUR_ACCESS_TOKEN
set GN_ACCESS_TOKEN_SECRET=YOUR_ACCESS_TOKEN_SECRET
```
for OSX
```
export GN_CONSUMER_KEY=YOUR_CONSUMER_KEY
export GN_CONSUMER_SECRET=YOUR_CONSUMER_SECRET
export GN_ACCESS_TOKEN=YOUR_ACCESS_TOKEN
export GN_ACCESS_TOKEN_SECRET=YOUR_ACCESS_TOKEN_SECRET
```
### Or edit Consumer Key and etc of main.go
If set env, always takes values form env.
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

### Binary
**Must** set env before execute.  
[Windows](https://dl.dropboxusercontent.com/u/35247301/go-nowplaying.exe)
## License
### MIT
