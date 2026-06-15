## Specification for component: checkBasicMIWU

Interface (in Go):

```go
func checkBasicMIWU(miwu data.MIWU) error
```

Return an Error if the MIWU is not safisfying one of the following conditions:

- The MIWU must have a name.
- The due date of the MIWU must not be in the past.
- The MIWU must have a set.
