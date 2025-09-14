# Monkey Language Interpreter

A **Go** implementation of the Monkey programming language interpreter, following the "Writing An Interpreter In Go" book by Thorsten Ball.

## Overview

This project implements the Monkey programming language as described in the book "Writing An Interpreter In Go" by Thorsten Ball. The Monkey language is a simple, dynamically typed programming language with first-class functions, closures, and a C-like syntax.

There is a **Rust** implementation of the Monkey language interpreter [here](https://github.com/rand0m42195/monkey-rs).

## Current Features

- **Lexer (Tokenizer)**: Converts source code into tokens
- **REPL (Read-Eval-Print Loop)**: Interactive command-line interface
- **Token Support**: 
  - Identifiers and integers
  - Arithmetic operators (`+`, `-`, `*`, `/`)
  - Comparison operators (`==`, `!=`, `<`, `>`)
  - Logical operators (`!`)
  - Delimiters (`(`, `)`, `{`, `}`, `,`, `;`)
  - Keywords (`let`, `fn`, `if`, `else`, `return`, `true`, `false`)

## Project Structure

```
monky-language/
├── go.mod              # Go module definition
├── main.go             # Entry point
├── token/              # Token definitions and keyword lookup
│   └── token.go
├── lexer/              # Lexical analysis
│   ├── lexer.go
│   └── lexer_test.go
└── repl/               # Interactive REPL
    └── repl.go
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd monky-language
```

2. Ensure you have Go 1.24.4 or later installed

3. Run the program:
```bash
go run main.go
```

## Usage

Start the Monkey REPL by running the program:

```bash
go run main.go
```

You'll see a welcome message and a prompt (`>>`). You can then type Monkey language code and see the tokenized output.

### Example Session

```
Hello username!, This is the Monkey programming language, version 0.0.1
Feel free to type in commands
>> let five = 5;
{Type:LET Literal:let}
{Type:IDENT Literal:five}
{Type:= Literal:=}
{Type:INT Literal:5}
{Type:; Literal:;}

>> let add = fn(x, y) { x + y; };
{Type:LET Literal:let}
{Type:IDENT Literal:add}
{Type:= Literal:=}
{Type:FUNCTION Literal:fn}
{Type:( Literal:(}
{Type:IDENT Literal:x}
{Type:, Literal:,}
{Type:IDENT Literal:y}
{Type:) Literal:)}
{Type:{ Literal:{}
{Type:IDENT Literal:x}
{Type:+ Literal:+}
{Type:IDENT Literal:y}
{Type:; Literal:;}
{Type:} Literal:}}
{Type:; Literal:;}

>> 10 == 10;
{Type:INT Literal:10}
{Type:== Literal:==}
{Type:INT Literal:10}
{Type:; Literal:;}
```

## Language Syntax

The Monkey language supports:

- **Variables**: `let name = value;`
- **Functions**: `fn(x, y) { x + y; }`
- **Conditionals**: `if (condition) { ... } else { ... }`
- **Returns**: `return value;`
- **Integers**: `5`, `10`, `42`
- **Booleans**: `true`, `false`
- **Arithmetic**: `+`, `-`, `*`, `/`
- **Comparison**: `==`, `!=`, `<`, `>`
- **Logical**: `!`

## Testing

Run the test suite:

```bash
go test ./...
```

## Development Status

This is an early-stage implementation with:
- ✅ Lexer (tokenizer) - Complete
- ✅ REPL interface - Complete
- 🚧 Parser - Not yet implemented
- 🚧 Evaluator - Not yet implemented
- 🚧 AST (Abstract Syntax Tree) - Not yet implemented

## Future Roadmap

- [ ] Implement parser to build Abstract Syntax Tree (AST)
- [ ] Add evaluator to execute the AST
- [ ] Support for more data types (strings, arrays, hashes)
- [ ] Error handling and reporting
- [ ] File execution (not just REPL)
- [ ] Standard library functions

## Related Projects

- **[RMonkey Language Interpreter (Rust)](https://github.com/rand0m42195/monkey-rs)**: A Rust implementation of the Monkey language interpreter with comprehensive features including lexer, parser, evaluator, and advanced language constructs.

## Language Comparison

This Go implementation focuses on simplicity and readability, following the original book's approach. The [Rust version](https://github.com/rand0m42195/monkey-rs) demonstrates more advanced features and showcases the differences between Go and Rust for interpreter development:

| Feature | Go Implementation | Rust Implementation |
|---------|------------------|-------------------|
| **Memory Management** | Garbage collected | Ownership system |
| **Performance** | Good | Excellent |
| **Error Handling** | Simple, explicit | Result<T, E> types |
| **Concurrency** | Goroutines | async/await |
| **Type System** | Simple, practical | Advanced, zero-cost abstractions |

## Contributing

This project is based on the "Writing An Interpreter In Go" book. Feel free to contribute improvements, bug fixes, or additional features.

When contributing, consider how your changes might compare to the [Rust implementation](https://github.com/rand0m42195/monkey-rs) and whether they maintain the Go version's focus on simplicity and educational value.

## License

This project is for educational purposes. Please refer to the original book for licensing information.

## Acknowledgments

- **Thorsten Ball** for the original "Writing An Interpreter In Go" book
- The Go community for excellent tooling and documentation
