#Mac OS
xk6 build --with xk6-fabric=. --with github.com/mostafa/xk6-kafka --output k6-darwin
#Linux
GOOS=linux GOARCH=amd64 xk6 build --with xk6-fabric=. --with github.com/mostafa/xk6-kafka --output k6-linux
