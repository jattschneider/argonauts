# Go Argonauts

Argonauts is a Go library that implements argo2 password hash.

## Getting Started

Just a quick example how to use the library:

#### main.go
```
package main

import (
	"fmt"

	"github.com/jattschneider/argonauts"
)

func main() {
	salt, err := argonauts.Salt()
	if err != nil {
		panic(err)
	}
	saltString := argonauts.Sprint(salt)
	fmt.Println(saltString)

	opts := argonauts.DefaultOptions(salt)
	passwd := "somerandompassword"
	hash := argonauts.Hash(opts, []byte(passwd))

	hashString := argonauts.Sprint(hash)
	fmt.Println(hashString)
	h, err := argonauts.ReadString(hashString)
	if err != nil {
		panic(err)
	}

	if match, err := argonauts.Compare(opts, []byte(passwd), h); err == nil && match {
		fmt.Println("password match!")
	}
}

```

```
$ go run main.go
```

### Installing

```
go get -v github.com/jattschneider/argonauts
```

## Built With

* [Go](https://golang.org/) - The Go Programming Language

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jattschneider/argonauts/tags). 

## Authors

* **José Augusto Schneider** - *Initial work* - [jattschneider](https://github.com/jattschneider)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
