# Presentations created with the Go present tool

Install the present tool:
- `go get golang.org/x/tools/present`
- cd into the `present` subdirectory
- Install `present` tool with `go install ./...`

Start presentation in this directory:
```
present -base . -notes
```

Get rendered HTML of presentation:
```
curl http://127.0.0.1:3999/.../present.slide > present.html
```
