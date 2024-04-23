# URDU Programming Language

It is simple compiler that compiles a dialect of BASIC to C, while being written in Go Lang

It supports:

- Numerical variables
- Basic arithmetic
- If statements
- While loops
- Print text and numbers
- Input numbers
- Labels and goto
- Comments

## How to run

```
go build -o compiler
./compiler <filename>.urdu
```

## Hello World

```
DEKHAO "Hello Dunya"

```

### Average of numbers program
```
NAM a = 0
JAB a < 1 KARO
    DEKHAO "Enter number of scores: "
    BTAO a
JABBND

NAM b = 0
NAM s = 0
DEKHAO "Enter one value at a time: "
JAB b < a KARO
    BTAO c
    NAM s = s + c
    NAM b = b + 1
JABBND

DEKHAO "Average: "
DEKHAO s / a
```

## Resources

Read the tutorial: [Let's make a Teeny Tiny compiler, part 1](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler1.html) as well as [part 2](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler2.html) and [part 3](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler3.html)
