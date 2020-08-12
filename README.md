# Usage

```go
a := struct {
    Bool   *bool
    Int    *int
    String *string
}{
    Bool:   ptr.Bool(true),
    Int:    ptr.Int(33),
    String: ptr.String("foo"),
}
```

```go
fn := func() (string) { return "value" }

var optionalString *string
if someCondition := true; someCondition {
    *ptr.NewString(&optionalString) = fn()
}
```
