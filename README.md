# Presentations created with the Go present tool

The present tool can be found here: `go get golang.org/x/tools/present`

Start presentation in this directory:
```sh
./bin/present.$(go env GOOS) -base . -notes
```

Get rendered HTML of presentation:
```sh
curl http://127.0.0.1:3999/.../present.slide > present.html
```

Convert a multislide version of a presentation to a normal one:
```sh
go run github.com/flowdev/present-go/cmd/multislide
```
