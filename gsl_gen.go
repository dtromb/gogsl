// GPL-2 License

package gogsl

// generate GSL bindings and pass the generated code thru gofmt.

//go:generate go run ./generator/generator.go
//go:generate gofmt -w .
