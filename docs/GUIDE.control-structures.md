# Language Guide: Control Structures

## if

`if` can be called in two different ways:

1. Method If: `conditional -> if: { true } { false }`
2. Function If: `if: { conditional } { true } { false }`

The conditional is evaluated based on the output produced by the
function and the exit number. Any non-zero exit numbers are an automatic
"false". Any functions returning no data are also classed as a "false".
For a full list of conditions that are evaluated to determine a true or
false state of a function, please read the documentation on the `boolean`
data type in [GUIDE.syntax](GUIDE.syntax.md#boolean).

Please also note that while the last parameter is optional, if it is
left off and `if` or `!if` would have otherwise called it, then `if` /
`!if` will return a non-zero exit number. The significance of this is
important when using `if` or `!if` inside a `try` block.

* Method If

This is where the conditional is evaluated from the result of the piped
function. The last parameter is optional.

    # if / then
    out: hello world | grep: world -> if: { out: world found }

    # if / then / else
    out: hello world | grep: world -> if: { out: world found } { out: world missing }

    if / else
    out: hello world | grep: world -> !if: { out: world missing }

* Function If

This is where the conditional is evaluated from the first parameter. The
last parameter is optional.

    # if / then / else
    if: { out: hello world | grep: world } { out: world found }

    # if / then / else
    if: { out: hello world | grep: world } { out: world found } { out: world missing }

    if / else
    !if: { out: hello world | grep: world } { out: world missing }


## !if

`if` also supports an anti-alias which will "not" the conditional,
effectively reversing the "then" and "else" parameters. See `if` (above)
for examples.

## try

This will force a different execution behaviour. All pipelined processes
will become sequential (unlike normally when they run in parallel) and
any exit numbers not equal to zero (0) will terminate the code block.
This also includes `if` statements so be very careful to include an else
parameter, even if it's an empty block, so `if` doesn't raise an error.

If the try block fails then try will raise a non-zero exit number. If
you want to run an alternative block of code in an event of a failure
then combine with the `catch` method.

    # try
    try: { out: hello world | grep: foobar; out: other stuff }

    # try / catch
    try: { out: hello world | grep: foobar; out: other stuff } \
        -> catch { out: `try` failed }

## catch

This works a little like the single parameter `!if` method except it
only checks the exit number (not stdin) and the stdin stream is simply
forwarded along the chain.

`catch` is typically used alongside `try` but it can also be used on its
own where you want to check the success of a routine while preserving
its stdout stream.

Use `!catch` to "else" the `try`

    # try / catch
    try: { out: hello world | grep: foobar; out: other stuff } -> catch { out: `try` failed }

    # catch
    out: hello world | grep: foobar -> catch { out: foobar not found }

    # !catch
    out: hello world | grep: world -> !catch { out: world found }

    # else
    try: { out: hello world | grep: foobar; out: other stuff } \
        -> catch  { out: `try` failed } \
        -> !catch { out: `try` succeeded }

`catch` also supports anti-alias (`!catch`) where the code block only
executes if the exit number equals zero. This is useful if you have a
block of code and want to report back a message or process additional
commands if the original block was successful.

## for

I was a little naughty in creating this builtin as it breaks one of my
style guidelines where the first parameter is input as a code block but
is actually treated as a string which is then split into 3 strings via
the semi-colon character, `;`, and processed as 3 separate blocks.

Usage:

    for ( i=1; i<6; i++ ) { echo $i }

The parameter is: `( i=1; i<6; i++ )`, but it is then converted into the
following code:

1. `let i=0` - declare the loop iteration variable
2. `= i<0` - if the condition is true then proceed to run the code in
the second parameter - `{ echo $i }`
3. `let i++` - increment the loop iteration variable

The reason behind breaking my own style guidelines on this function was
to create a for loop that was mirrored what the average developer might
expect. However `foreach` and `while` both follow the usual structure of
_murex_ code or want the performance gained from not running 3 additional
commands on each loop iteration, then you could write it using the array
function, `a`:

    a: [1..5] -> foreach i { echo $i }

Using the `a` function like this is more idiomatic _murex_ as well as
much more performant. For example, in the commands below, both print
every multiple of 2 between 10 and 10,000 (inclusive). The former is
less readable but will take approximately 1/10th of the time to process
when compared to the latter:

    time { a: [1..999][0,2,4,6,8],10000 -> foreach i { $i } }
    time { for { i=10; i<10001; i=i+2 } { $i } }

For more information about the `array` function see [GUIDE.arrays-and-maps](GUIDE.arrays-and-maps.md#the-array-builtin).

## foreach / formap

These functions will loop though each element in an array (`foreach`) or
map (`formap`)

    listed_output -> foreach: var_name { iteration }
    structured_output -> formap: var_key var_value { iteration }

More detailed descriptions on working with structured data, including
examples, can be found in [GUIDE.arrays-and-maps](GUIDE.arrays-and-maps.md#working-with-structured-data).

## while
(description to follow)

    while: { conditional } { iteration }

    select: { condition } {
