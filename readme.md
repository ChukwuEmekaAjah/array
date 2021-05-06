<h1 align="center">array</h1>

A higher order wrapper over Go slices that allows use of methods similar to JavaScript arrays higher-order methods

## What you get
- Use simplified methods on arrays without reinventing the wheel every time.

## How to use
```bash
go get github.com/ChukwuEmekaAjah/array
```

```go

    package main

    import (
        "errors"
        "fmt"

        "github.com/ChukwuemekaAjah/array"
    )

    func main() {
        items := []int{23, 24, 2, 5, 10}
	    interfaceItems := make([]interface{}, len(items))

        for i, v := range items {
            interfaceItems[i] = v
        }

        a := array.New(interfaceItems)

        fmt.Println("length of array", a.Length())

        fmt.Println("some are even", a.Some(isEven))

        fmt.Println("all are even", a.Every(isEven))
    }

    func isEven(value interface{}, index int) bool {

        if value.(int) % 2 {
            return false
        }

        return true
    }

    
```

## API 
Methods on the <b>Array</b> struct are:
- array.New()
- array.Filter()
- array.Map()
- array.ForEach()
- array.Some()
- array.Length()
- array.Every()

## Understanding API methods and properties
WIP


## Contributing
In case you have any ideas, features you would like to be included or any bug fixes, you can send a PR.

- Clone the repo

```bash
git clone https://github.com/ChukwuEmekaAjah/array.git
```

## Todo
- Write more documentation on the explanation of array methods parameters.