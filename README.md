# Monke(Monkey) Interpreter 🐵

**Monkey(Monkey) Interpreter**, a simple programming language interpreter written in **Go**, based on the book *["How to Write an Interpreter in Go"](https://interpreterbook.com/)* by **Thorsten Ball**.

This project implements the **Monkey** language—a dynamically typed, Lisp-like language—featuring first-class functions, closures, and a simple, expressive syntax. The Monkey Interpreter is built from scratch and showcases core concepts of language design, parsing, and evaluation.

## Features

- **Arithmetic Operations**: `+`, `-`, `*`, `/`
- **Booleans & Logic**: `true`, `false`, `!`, `&&`, `||`
- **Comparisons**: `==`, `!=`, `>`, `<`, `>=`, `<=`
- **Variables**: Declaration and reassignment with `let`z
- **Conditionals**: `if`-`else` for branching logic
- **First-Class Functions**: Pass, return, and assign functions to variables
- **Closures**: Functions capture and use variables from their surrounding scope
- **Arrays**: Dynamic arrays with indexing
- **Hashmaps**: Key-value pairs with flexible key types
- **Built-in Functions**: Standard functions like `len`, `first`, `last`, `push`
- **REPL**: A Read-Eval-Print-Loop for live interaction with Monkey

## Getting Started

### Prerequisites

- **Go 1.18+**: Install Go from [golang.org](https://golang.org/dl/).

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/your-username/monkey-interpreter.git
    cd monkey-interpreter
    ```

2. Build the interpreter:

    ```bash
    go build -o monkey
    ```

### Running the REPL

To start the Monkey language REPL (Read-Eval-Print-Loop):

```bash
./monkey
```

### Possible Future Work 🛠️

- **Improve Error Handling**: Better error messages for easier debugging.
- **Performance Enhancements**: Optimize the evaluator for faster execution.
- **New Features**: Add features like loops, more data structures, or additional built-in functions.# MonkeLang-Interpreter
# MonkeLang-Interpreter
