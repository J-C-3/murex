# Language Guide: Quick start guide for Bash developers

This is a cheat sheet reference for lazy Bash developers wishing to
accelerate their introduction into _murex_:

## Piping and redirection

_murex_ supports the `|` pipe just like Bash but the preferred pipe
token in _murex_ the arrow, `->` (those two token are interchangeable).

Redirection of stdout and stderr is very different in _murex_. There is
no support for the `2>` or `&1` tokens,  instead you name the pipe as
the first parameter:

    err: <!out> "error message redirected to stdout"
    out: <err> "output redirected to stderr"

You can also use named pipes this way to join up parts of the script
that otherwise wouldn't be part of the same pipeline. See [GUIDE.syntax](GUIDE.syntax.md#piping)
for more details on named pipes.

To redirect to a file you can use the `>` or `>>` functions. They work
similarly to bash except that they are functions rather than tokens. This means
they literally work like the following:

    out: "message" -> >  new-file.txt
    out: "message" -> >> append-file.txt

However this is clearly ugly in practice. So the following syntactic sugar is
supported, `|>` for overwrite and `>>` for append:

    out: "message" |> new-file.txt
    out: "message" >> append-file.txt

## Emendable sub-shells

There are two types of emendable sub-shells: strings and arrays.

* string sub-shells, `${ command }`, take the results from the sub-shell
and return it as a single parameter. Equivalent to the following in bash:
`command "$(sub-shell command)"`.

* array sub-shells, `@{ command }`, take the results from the sub-shell
and expand it as parameters. Arrays can be multiple lines (like in Bash)
or array objects in more complex data formats like JSON. Unlike bash,
other white spaces such as tabs and space characters are not counted as
separators for walking through arrays. This is intentional to allow line
formatting and space characters in file names. Array shells are
equivalent to the following in Bash: `command $(sub-shell command)`

Examples:

    ls -l ${out: file name}           # works because file name contain space
    ls -l @{out: file1 file2 file3}   # fails because not an array
    ls -l @{out: file1\nfile2\nfile3} # works because output is an array

The reason _murex_ breaks from the POSIX tradition of using backticks and
parentheses is because _murex_ works on the principle that everything inside
a curly bracket is considered a new block of code. Typically that would mean
a sub-shell however sometimes it could be configuration code in the form of
inlined JSON.

## Globbing

There isn't auto-expansion of globbing to protect against accidental
damage. Instead globbing is achieved via sub-shells using either:

* `g` (traditional globbing)
* `rx` (regexp matching in current directory only)
* `f` (file or directory type matching)

Examples:

    # all text files via globbing:
    ls -l @{g *.txt}

    # all text and markdown files via regexp:
    ls -l @{rx '\.(txt|md)$'}

    # all files via type matching:
    ls -l @{f +f}

You can also using type matching against globbing and regexp to filter
out types in conjunction with file name matching:

    # all directories named *.txt
    ls -l @{g *.txt -> f +d}

## Exit code

In bash the variable `$?` would store the exit code. This doesn't exist
in _murex_. Instead there a separate command `exitnum`:

    open: test/fox.txt -> grep: foobar; exitnum

## Array expansion

In bash you can expand arrays using the following syntax: `a{1..5}b`. In
_murex_ this is another sub-shell process: `a: a[1..5]b`. As you can see,
_murex_ also uses square brackets instead as well. There are a few other
changes, read [GUIDE.arrays-and-maps](GUIDE.arrays-and-maps.md#the-array-builtin)
for more on using the `array` builtin.

## Back ticks

In _murex_ back ticks do not spawn sub-shells. Back ticks are treated
like a regular, printable, character. Their only special function is
quoting strings in `=`, eg:

    if { = `quoted string`==variable } { out "do something" }
