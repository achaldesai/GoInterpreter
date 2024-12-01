# GoInterpreter

GoInterpreter is a simple interpreter made with the Go programming language.
This project aims to provide a basic understanding of how interpreters work
and to serve as a learning tool for those interested in programming language
design and implementation.

## Features

- Basic syntax parsing
- Variable declaration and assignment
- Arithmetic operations
- Conditional statements
- Loop constructs

## Installation

To install GoInterpreter, clone the repository and build the project:

```bash
git clone https://github.com/achaldesai/GoInterpreter.git
cd GoInterpreter
go build
```

## Usage

To run the interpreter, use the following command:

```bash
./GoInt
```

## Examples

Here are some example lines that you can run with GoInterpreter:

```bash
let a = 5; a;
let add = fn(x, y) { x + y; }; add(5, 5);
len("hello world")
let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];
let two = "two";
let testHash = {
"one": 10 - 9,
two: 1 + 1,
"thr" + "ee": 6 / 2,
4: 4,
true: 5,
false: 6
}
testHash;
testHash["three"]
testHash[true]
```

## Acknowledgements

- [Go Programming Language](https://golang.org/)
- [Writing An Interpreter In Go](https://interpreterbook.com/)
