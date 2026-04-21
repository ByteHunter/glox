# About

This project is an implementation of an interpreter for the fictional programming language Lox from the book "Crafting Interpreters", by Robert Nystrom.

This implementation is written in Go.

# Grammar

```
progam              -> declaration* EOF ;

declaration         -> variableDeclaration | statement ;

variableDeclaration -> "var" IDENTIFIER ( "=" expression )? ";" ;

statement           -> expressionStatement | printStatement ;

expressionStatement -> expression ";" ;
printStatement      -> "print" expression ";" ;

expression          -> assignment ;
assignment          -> IDENTIFIER "=" assignment | equality ;
equality            -> comparison ( ( "!=" | "==" ) comparison )* ;
comparison          -> term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term                -> factor ( ( "-" | "+" ) factor )* ;
factor              -> unary ( ( "/" | "*" ) unary )* ;
unary               -> ( "!" | "-" ) unary | primary ;
primary             -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" | IDENTIFIER ;
```
