# Znote

Note compiling tool based on pandoc.

## About

Compiles documents from pandoc markdown format to PDF format creating a table of contents and including a few preprocess directives.

## Dependencies

* Pandoc

## Installation

Copy znote to somewhere in your PATH.

## Usuage

To initialize a new note directory create a new directory and run `znote -init`

You will then need to make sure your profile has `eval $(znote -alias)` in it.

To create a new note file simply type nnotes

## Features

* All pandoc markdown syntax
* LaTeX Math preprocessor directives


## Preprocessor directives

Before compilation from Markdown to PDF znote will expand a few functions to be full LaTeX functions.

Start a line with `=>` to have it be converted to a right arrow for math equation continuation. Example

```
  $a^2+b^2=c^2$\
  $=> \sqrt{a^2 + b^2}$
```

Expands to

```
  $a^2+b^2=c^2$\
  $\Rightarrow \sqrt{a^2 + b^2}$
```

This is a very simple example but I have found saves a lot of time on equations that use 6+ lines to show work.

Shortened infinity syntax from `\infty` to `\inf`, very straightforward.

```
  $\lim_{n\to\inf}$
```
Expands to

```
  $\lim_{n\to\infty}$
```

Just added this to save a few keystrokes each time I want to insert infinity.

Intgral and Summation limit reposition. Moves the limits of either integration or summation from the right side to the top and bottom of the symbole, only affects small math text(wrapped in single $)

```
  $\int_{1}^{\inf} f(x)\,dx$
```
Expands to
```
  $\int\limits_{1}^{\inf} f(x)\,dx$
```

This just makes in-line integrals and summations easier to read.





