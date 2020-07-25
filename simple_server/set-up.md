## server.go

Simple boilerplate for go server. For testing, try using telnet. To install:

``` 
brew install telnet
```

To start the server, navigate to your working directory where 'server.go` is located and run:

``` 
➜  simple_server go run server.go
Running server. Listening at port 8080...
```

To specify port number:
``` 
➜  simple_server go run server.go 9000
Running server. Listening at port 9000...
```



To connect to your server, go to your terminal/shell environment:

```
➜  ~ telnet localhost 8080
Trying ::1...
Connected to localhost.
Escape character is '^]'.

You've connected to the server at localhost:8080
Connection open.........
Closing connection......
Connection closed by foreign host.
➜  ~
```
