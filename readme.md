# Simple Compiler wriiten in Go

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

## Example Code

```
PRINT "How many fibonacci numbers do you want?"
INPUT nums
PRINT ""

LET a = 0
LET b = 1
WHILE nums > 0 REPEAT
    PRINT a
    LET c = a + b
    LET a = b
    LET b = c
    LET nums = nums - 1
ENDWHILE
```

## Resources

It is basically rewrite of [teenytinycompiler](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler1.html) by [AZHenley](https://github.com/AZHenley) by that is written in python
Read the tutorial: [Let's make a Teeny Tiny compiler, part 1](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler1.html) as well as [part 2](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler2.html) and [part 3](http://web.eecs.utk.edu/~azh/blog/teenytinycompiler3.html)
