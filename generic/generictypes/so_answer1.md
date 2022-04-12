# Instantiating generically typed variables in Go

### Response to Stack Overflow question [71677581](soPost)

>--Michael Treanor  [GitHub](github)  [Twitter](twitter)
---

*Code for this post is available on:*

[Go Playground](playground) or [GitHub](repo)

MIT License

---
## Instantiate

First, to fix what you have, you must ***"instantiate"*** the new variable `mapping` of the given type `IExample` to specify which of the valid type constraints is being implemented in this instance of the type.

This means that one of the allowed types is chosen and 'set in stone' at compile time, allowing the compiler to do static type checking and avoiding the complexity and performance loss of the reflect package.

>In this example, since the type ***constraint*** chosen is `any`, (an alias of `interface{}`, the instantiated type may be any valid type.

## ***TLDR***:
***The compiler creates the correct "version" of the type for the situation you want while it is compiling.***

>In this example, you define a variable of the type `IExample`. Since `IExample` is an `interface`, any variable of that has a method `ExampleFunc(ex T) T` will work.

```go
// IExample is the original type from the question.
type IExample[T any] interface {
    ExampleFunc(ex T) T
}

var mapping map[string]IExample[any]
```
---
## Create variables
Create variables of type `IExample` to use as values in the map. Here are a few examples that instantiate `IExample` for types `any` and `int`.

You may notice that using an interface type is redundant and leads to a noticeable increase in the clutter since any instantiated type must have an ExampleFunc method. This is the major problem that generics are expected to solve: having to type out a method for each type.

Functions, channels, or other objects may be better suited. For now, I'll keep with your example.
```go
// anyThing is a sample of some object
// that implements IExample
type anyThing struct{ any }

func (t anyThing) ExampleFunc(ex any) any             { return t.any }

// intThing is a similar object that
// impements IExample ...
type intThing struct{ int }

func (t intThing) ExampleFunc(ex int) int             { return t.int }
```
---
## Assign keys and values
Load the *keys* and *values* into your `map`.
* Note that values of intThing (above) will not work since the map is instantiated with type `any`.
```go
// mapping maps string names to examples. The
// examples are instantiated with [any] and
// thus are filled in with interface objects
// that implement ExampleFunc(ex any) any
var mapping = map[string]IExample[any]{
    "any1": anyThing{"stuff"},
    "any2": anyThing{"other stuff"},
}

// add another example
mapping["any3"] = anyThing{"more different stuff"}

```
## Print results to test or run Example tests if you have them.
```go

func ExampleMapping() {
    for k, v := range mapping {
        fmt.Printf("%v: %v\n", k, v)
    }

    // Output:
    // any1: {stuff}
    // any2: {other stuff}
    // any3: {more different stuff}
}

```
---
## Complie time error
>Try to add an example with a different type constraint and see what happens ...

* Because the `mapping` is already instantiated to type `any`, you cannot add values of other types. For example:

```go
mapping["float64"] = float64thing{42.0}
// compile time error: InvalidIfaceAssign
```
```go
/*************  InvalidIfaceAssign  *************
cannot use (float64thing literal) (value of type float64thing) as IExample[any] value in map literal: float64thing does not implement IExample[any] (wrong type for method ExampleFunc)
 		have ExampleFunc(ex float64) float64
 		want ExampleFunc(ex any) any m
************************************************/
```

* The mapping variable we instantiated with type `any` will work with float64 values, but another can be instantiated to do so:
```go
var mappingFloat map[string]IExample[float64]
```
---
## Other choices

### Specific Name
You may prefer to instantiate the type with a new name so it can be used without the type arguments:

```go
type IExampleAny = IExample[any]
type IExampleFloat = IExample[float64]

var mapping map[string]IExampleAny
var mappingFloat map[string]IExampleFloat
```
### Single Type instead of conglomeration
You may prefer to build the type into a one-piece IExampleMap instead of a separate interface and map:

```go
type IExampleMap[T any] map[string]interface{ ExampleFunc(ex T) T }

var mapping IExampleMap[any]
```
or, even better:

```go
type IExampleMapAny = IExampleMap[any]

var mapping IExampleMapAny
```

### Use generics ... as generics (avoid interfaces)
Or you may even wamt to do away with the interface entirely and simply store the function signature. This way, it can be named anything, instead of being locked into implementing ExampleFunc(ex T) T. This allows greater flexibility.

Since the interface in this example implements no other methods, I am assuming its only feature would be to call that function ... so skip all that and just call the function.

One of the major reasons that generics have been the top request of Go developers for over 10 years is that they greatly reduce the boilerplate and testing requirements for separate functions and methods whose only difference is they type that they handle.

In addition, it has been made clear that using generics in Go is already much faster than reflect and interface{} even in the beta version of 1.18. Reflect is a fun and interesting package, but it was not designed for performant regular production applications.

Using a function instead of an interface in your example also makes it easy to simply use a basic function that is not a method on any object at all ... a wrapper function ... a closure ... or an inline function ... even a go func() ...

And ... it still works with your interfaces just like it always did ... because they have the same function signature.

```go
type Example[T any] func(ex T) T
//   - or -
type ExampleMap[T any] map[string]func(ex T) T

// instantiated with int:
var intMap ExampleMap[int]

func MyFuncInt(ex int) int { return 42 }
func AddOne(ex int) int { return x + 1 }

intMap["MyFunc"] = MyFuncInt
intMap["AddOne"] = AddOne
```
Another Example:
```go

func MyFuncAny(ex any) any { return "stuff" }

// Mapping2 maps string names to examples. The examples
// are stored as functions with the signature
// 		func(ex T) T
// which is instantiated to
// 		func(ex any) any
// in this example
var Mapping2 = ExampleMap[any]{

    // custom function
    "any": MyFuncAny,

    // original interface
    "a": anyThing{"stuff"}.ExampleFunc,

    // inline function
    "fake": func(ex any) any{return "fake"},
}

func ExampleMapping2() {
    for k, v := range Mapping2 {
        // calling v with the zero value of the type for this
        // example; this data could be stored in a struct,
        // sent on a channel, stored in a file ...
		fmt.Printf("%v: %v\n", k, v(nil))
	}

    // Output:
    // any: stuff
    // a: stuff
    // fake: fake
}

```
[github]:(https://github.com/skeptycal)
[twitter]:(https://twitter.com/skeptycal)
[soPost]: (https://stackoverflow.com/q/71677581)
[playground]: (https://go.dev/play/p/jR5MpX-Fopc)
[repo]: (https://github.com/skeptycal/71677581)