# Presentation About the Pros and Cons of Event Sourcing
Show the do's and don'ts of event sourcing and alternatives where appropriate.

Install the present tool:
- `go get golang.org/x/tools/present`
- cd into the `present` subdirectory
- Install `present` tool with `go install ./...`

Start presentation:
```
present -base . -notes
```

Get rendered HTML of presentation:
```
curl http://127.0.0.1:3999/present.slide > present.html
```
