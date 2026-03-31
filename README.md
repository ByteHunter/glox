# About

This project is an implementation of an interpreter for the fictional programming language Lox from the book "Crafting Interpreters", by Robert Nystrom.

This implementation is written in Go.

# Grammar

```
expression -> literal | unary | binary | grouping ;
literal    -> NUMBER | STRING | "true" | "false" | "nil" ;
grouping   -> "(" expression ")" ;
unary      -> ( "-" | "!" ) expression ;
binary     -> expression operator expression ;
operator   -> "==" | "!=" | "<" | "<=" | ">" | ">=" | "+" | "-" | "*" | "/" ;
```
