# upstruct [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/hackirby/upstruct)

upstruct is an utility package to update a golang struct with another one

## Install

```bash
go get github.com/hackirby/upstruct
```

## Example

```go
import "github.com/hackirby/upstruct"

type User struct {
    Username string
    Password string
}

type UserPatch struct {
    Username string
    Password string
}

var user = User{
    Username: "user",
    Password: "password",
}

var userPatch = UserPatch{
    Username: "newuser",
    Password: "newpassword",
}

func main() {
    upstruct.Update(&user, userPatch)
}
```
