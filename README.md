# go-libregf
Go bindings for the libregf C library, which exposes an API for handling Windows Registry files

# What is Working

* registry files (open, root key, get key, get value)
* keys (name, classname, values, subkeys)
* values (name, value, support for most types)
* error handling

# How to Use

```go
inport (
  "fmt"
  "github.com/jdrowell/go-libregf"
)

func report(filepath string, values []string) error {
  file, err := libregf.OpenFile(filepath)
  if err != nil { return err }
  defer file.Close()

  for _, path := range values {
    v, err := file.Value(path)
    if err != nil { return err }
    fmt.Printf("%s: %s\n", path, v)
  }

  return nil
}

func main() {
  err := report("SOFTWARE", []string{
    "Microsoft\\Windows NT\\CurrentVersion\\RegisteredOrganization",
    "Microsoft\\Windows NT\\CurrentVersion\\RegisteredOwner",
    "Microsoft\\Windows NT\\CurrentVersion\\ProductName",
    "Microsoft\\Windows NT\\CurrentVersion\\InstallDate",
    "Microsoft\\Windows NT\\CurrentVersion\\ProductId",
  })
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
```

# Dependencies

These bindings are up to date with release alpha-20230319 of [libregf](https://github.com/libyal/libregf).
You must have libregf-dev installed to be able to link your binary. Also, make sure **NOT** to have something
like <code>CGO_ENABLED=0</code> in your environment.

# Documentation

There is inline documentation in <code>go doc</code> format. Just use that command to explore it.
