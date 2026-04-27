# Bridge
![GitHub License](https://img.shields.io/github/license/The-Bridge-Foundation/BridgeToolchain?style=plastic)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/The-Bridge-Foundation/BridgeToolchain?style=plastic)

**Bridge** is a general-purpose programming language, designed to be easily embedded in other programs. The syntax follows a C-like philosophy to make the code easy to read.

## Overview

Bridge is a strongly-typed, compiled language with a clean syntax targeting its own VM. The toolchain consists of:

- **`bridgec`**: the Bridge compiler
- **`bridger`**: the Bridge runtime/linker
- **`bridgelib`**: the Bridge library

## Dependencies

- Golang 1.25
- LLVM 19

## Building

On UNIX-like systems, you can just use the build script:
```sh
./build.sh
```

## Install

On UNIX-like systems, you can just use the install script:
```sh
./install.sh
```

This installs `bridgec` and `bridger` to `~/.bridge/bin/` and adds them to your `PATH` via your shell config. The install script also configures `$BRIDGEPATH` for package retrieval.

## Usage

```
bridgec [flags] <source.bridge>

Flags:
  --version      Print version info
  --versionjson  Print version info as JSON
  -q             Quiet mode (suppress output)
  --dumpllvm     Print LLVM IR to stdout
  --printast     Print all AST nodes
```

The compiler outputs a `.bb` binary object, which can be run with bridger.

```
bridger [flags] <file.bb>

Flags:
  --version      Print version info
  --versionjson  Print version info as JSON
  -q             Quiet mode (suppress output)
  --maxmemory    Max pages allocable by the VM
```

## Language

### Hello World

```bridge
@Import(Console)

func Main(args string[]) {
    Console::Print("Hello, world!")
}
```

### Types

Primitive types: `int`, `float`, `string`, `bool`  
Nullable types: append `?` to any type (e.g. `int?`)

```bridge
set x int         // user-provided type
let y = 9         // inferred type
const z int = 9   // constant
```

### Custom Objects

```bridge
obj MyObject {
    foo int
    bar func(self, int) int
}

func MakeMyObject(a int) MyObject {
    set mobj MyObject
    mobj.foo = a
    mobj.bar = func (s self, c int) int {
        return c + s.foo
    }
    return mobj
}
```

### Control Flow

```bridge
loop { ... }            // loop forever (break to exit)
times 5 { ... }         // repeat N times
while condition { ... } // while loop
each x [0,1,2] { ... }  // iterate array elements
escape                  // break out of all loops
break                   // break out of the current loop
```

### Switch

```bridge
switch value {
    case 1: ...
    case 2: ...
    default: ...
}
```

### Arrays and Memory

```bridge
let array = [0, 0, 0]      // inline array
$varname                   // get address of variable
```

### typeof

```bridge
Console::Printf("%s\n", typeof x)  // prints the type name
```

### Annotations

| Annotation | Description |
|-----------|-------------|
| `@Import(Module)` | Import a module's exports into scope |
| `@Journaled()` | Export the following symbol |

### Namespace Access

```bridge
@Import(Console)
...
Console::Printf("%d\n", value)
Console::Print("text")
```

### Packages

In Bridge, you can make packages by marking the exported symbols with the @Journaled notation and filling out the spec.json file that will go alongside the bb executable. Once that's done, you can place the package in $BRIDGEPATH/packages and write code using it.

## Project Structure

| File | Description |
|------|-------------|
| `bridgec/` | Compiler |
| `bridgelib/` | VM library |
| `bridger/` | VM executable |
| `examples/` | Source code examples |
| `stdlib/` | Built-in standard library |

## License

Bridge is licensed under Apache License 2.0.

See [NOTICE](NOTICE) for further details.
