package docs

func init() {

	Definition["and"] = "# _murex_ Shell Docs\n\n## Command Reference: `and`\n\n> Returns `true` or `false` depending on whether multiple conditions are met\n\n### Description\n\nReturns a boolean results (`true` or `false`) depending on whether all of the\ncode-blocks included as parameters are successful or not.\n\n### Usage\n\n    and: { code-block } { code-block } -> <stdout>\n    \n    !and: { code-block } { code-block } -> <stdout>\n    \n`and` supports as many or as few code-blocks as you wish.\n\n### Examples\n\n    if { and { = 1+1==2 } { = 2+2==4 } { = 3+3==6 } } then {\n        out: The laws of mathematics still exist in this universe.\n    }\n\n### Detail\n\n`and` does not set the exit number on failure so it is safe to use inside a `try`\nor `trypipe` block.\n\nIf `and` is prefixed by a bang then it returns `true` only when all code-blocks\nare unsuccessful.\n\n#### Code-Block Testing\n\n* `and` tests all code-blocks up until one of the code-blocks is unsuccessful,\n  then `and` exits and returns `false`.\n\n* `!and` tests all code-blocks up until one of the code-blocks is successful,\n  then `!and` exits and returns `false` (ie `!and` is `not`ing every code-block).\n\n### Synonyms\n\n* `and`\n* `!and`\n\n\n### See Also\n\n* [commands/`!` (not)](../commands/not.md):\n  Reads the STDIN and exit number from previous process and not's it's condition\n* [commands/`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe` \n* [commands/`false`](../commands/false.md):\n  Returns a `false` value\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n* [commands/`true`](../commands/true.md):\n  Returns a `true` value\n* [commands/`try`](../commands/try.md):\n  Handles errors inside a block of code\n* [commands/`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error"

}
