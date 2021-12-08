# _murex_ Shell Docs

## Command Reference: `a` (mkarray)

> A sophisticated yet simple way to build an array or list

## Description

Pronounced "make array", like `mkdir` (etc), _murex_ has a pretty sophisticated
builtin for generating arrays. Think like bash's `{1..9}` syntax:

    a: [1..9]

## Usage

    a: [start..end] -> <stdout>
    a: [start..end.base] -> <stdout>
    a: [start..end,start..end] -> <stdout>
    a: [start..end][start..end] -> <stdout>
    
All usages also work with `ja` and `ta` as well:

    ja: [start..end] -> <stdout>
    ta: data-type [start..end] -> <stdout>

## Examples

    » a: [1..3]
    1
    2
    3
    
    » a: [3..1]
    3
    2
    1
    
    » a: [01..03]
    01
    02
    03

## Detail

### Alternative Number Bases

You can also specify an alternative number base by using an `x` or `.`
in the end range:

    a: [00..ffx16]
    a: [00..ff.16]
    
All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.
Please note that the start and end range are written in the target base
while the base identifier is written in decimal: `[hex..hex.dec]`

Also note that the additional zeros denotes padding (ie the results will
start at `00`, `01`, etc rather than `0`, `1`...)

### Character arrays

You can select a range of letters (a to z):

    » a: [a..z]
    » a: [z..a]
    » a: [A..Z]
    » a: [Z..A]
    
...or any characters within that range.

### Special ranges

Unlike bash, _murex_ also supports some special ranges:

```  
» a: [mon..sun]
» a: [monday..sunday]
» a: [jan..dec]
» a: [january..december]
» a: [spring..winter]
```

It is also case aware. If the ranges are uppercase then the return will
be uppercase. If the ranges are title case (capital first letter) then
the return will be in title case:

    » a: [Monday..Sunday]
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
    
Where the special ranges differ from a regular range is they cannot
cannot down. eg `a: [3..1]` would output

    » a: [3..1]
    3
    2
    1
    
however a negative range in special ranges will cycle through to the end
of the range and then loop back from the start:

    » a: [Thursday..Wednesday]
    Thursday
    Friday
    Saturday
    Sunday
    Monday
    Tuesday
    Wednesday
    
This decision was made because generally with ranges of this type, you
would more often prefer to cycle through values rather than iterate
backwards through the list.

If you did want to reverse then just pipe the output into another tool:

    » a: [Monday..Friday] -> mtac
    Friday
    Thursday
    Wednesday
    Tuesday
    Monday
    
There are other UNIX tools which aren't data type aware but would work in
this specific scenario:

* `tac` (Linux),

* `tail -r` (BSD / OS X)

* `perl -e "print reverse <>"` (Multi-platform but requires Perl installed)

### Advanced Array Syntax

The syntax for `a` is a comma separated list of parameters with expansions
stored in square brackets. You can have an expansion embedded inside a
parameter or as it's own parameter. Expansions can also have multiple
parameters.

    » a: 01,02,03,05,06,07
    01
    02
    03
    05
    06
    07
    
    » a: 0[1..3],0[5..7]
    01
    02
    03
    05
    06
    07
    
    » a: 0[1..3,5..7]
    01
    02
    03
    05
    06
    07
    
    » a: b[o,i]b
    bob
    bib
    
You can also have multiple expansion blocks in a single parameter:

    » a: a[1..3]b[5..7]
    a1b5
    a1b6
    a1b7
    a2b5
    a2b6
    a2b7
    a3b5
    a3b6
    a3b7
    
`a` will cycle through each iteration of the last expansion, moving itself
backwards through the string; behaving like an normal counter.

### Creating JSON arrays with `ja`

As you can see from the previous examples, `a` returns the array as a
list of strings. This is so you can stream excessively long arrays, for
example every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`
(this kind of array expansion would hang bash).

However if you needed a JSON string then you can use all the same syntax
as `a` but forgo the streaming capability:

    » ja: [Monday..Sunday]
    [
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
        "Sunday"
    ]
    
This is particularly useful if you are adding formatting that might break
under `a`'s formatting (which uses the `str` data type).

## See Also

* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [types/`str` (string) ](../types/str.md):
  string (primitive)
* [commands/`ta` (mkarray)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type