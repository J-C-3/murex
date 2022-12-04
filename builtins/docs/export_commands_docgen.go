package docs

func init() {

	Definition["export"] = "# _murex_ Shell Docs\n\n## Command Reference: `export`\n\n> Define an environmental variable and set it's value\n\n## Description\n\nDefines, updates or deallocates an environmental variable.\n\n## Usage\n\n    <stdin> -> export var_name\n    \n    export var_name=data\n\n## Examples\n\nAs a method:\n\n    » out \"Hello, world!\" -> export hw\n    » out \"$hw\"\n    Hello, World!\n    \nAs a function:\n\n    » export hw=\"Hello, world!\"\n    » out \"$hw\"\n    Hello, World!\n\n## Detail\n\n### Deallocation\n\nYou can unset variable names with the bang prefix:\n\n    !export var_name\n    \nFor compatibility with other shells, `unset` is also supported but it's really\nnot an idiomatic method of deallocation since it's name is misleading and\nsuggests it is a deallocator for local _murex_ variables defined via `set`.\n\n### Exporting a local or global variable\n\nYou can also export a local or global variable of the same name by specifying\nthat variable name without a following value. For example\n\n    # Create a local variable called 'foo':\n    » set: foo=bar\n    » env -> grep: foo\n    \n    # Export that local variable as an environmental variable:\n    » export: foo\n    » env -> grep: foo\n    foo=bar\n    \n    # Changing the value of the local variable doesn't alter the value of the environmental variable:\n    » set: foo=rab\n    » env -> grep: foo\n    foo=bar\n    » out: $foo\n    rab\n    \n### Scoping\n\nVariable scoping is simplified to three layers:\n\n1. Local variables (`set`, `!set`, `let`)\n2. Global variables (`global`, `!global`)\n3. Environmental variables (`export`, `!export`, `unset`)\n\nVariables are looked up in that order of too. For example a the following\ncode where `set` overrides both the global and environmental variable:\n\n    » set:    foobar=1\n    » global: foobar=2\n    » export: foobar=3\n    » out: $foobar\n    1\n    \n#### Local variables\n\nThese are defined via `set` and `let`. They're variables that are persistent\nacross any blocks within a function. Functions will typically be blocks\nencapsulated like so:\n\n    function example {\n        # variables scoped inside here\n    }\n    \n...or...\n\n    private example {\n        # variables scoped inside here\n    }\n    \n    \n...however dynamic autocompletes, events, unit tests and any blocks defined in\n`config` will also be triggered as functions.\n\nCode running inside any control flow or error handing structures will be\ntreated as part of the same part of the same scope as the parent function:\n\n    » function example {\n    »     try {\n    »         # set 'foobar' inside a `try` block\n    »         set: foobar=example\n    »     }\n    »     # 'foobar' exists outside of `try` because it is scoped to `function`\n    »     out: $foobar\n    » }\n    example\n    \nWhere this behavior might catch you out is with iteration blocks which create\nvariables, eg `for`, `foreach` and `formap`. Any variables created inside them\nare still shared with any code outside of those structures but still inside the\nfunction block.\n\nAny local variables are only available to that function. If a variable is\ndefined in a parent function that goes on to call child functions, then those\nlocal variables are not inherited but the child functions:\n\n    » function parent {\n    »     # set a local variable\n    »     set: foobar=example\n    »     child\n    » }\n    » \n    » function child {\n    »     # returns the `global` value, \"not set\", because the local `set` isn't inherited\n    »     out: $foobar\n    » }\n    » \n    » global: $foobar=\"not set\"\n    » parent\n    not set\n    \nIt's also worth remembering that any variable defined using `set` in the shells\nFID (ie in the interactive shell) is localised to structures running in the\ninteractive, REPL, shell and are not inherited by any called functions.\n\n#### Global variables\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill be scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\n#### Environmental variables\n\nExported variables (defined via `export`) are system environmental variables.\nInside _murex_ environmental variables behave much like `global` variables\nhowever their real purpose is passing data to external processes. For example\n`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):\n\n    » export foo=bar\n    » env -> grep foo\n    foo=bar\n    \n### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> set hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » set cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nThis only works for external executables. There is currently no way to call\naliases, functions nor builtins from a variable and even the above `exec` trick\nis considered bad form because it reduces the readability of your shell scripts.\n\n### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n    » out \"$foo\"\n    bar\n    \n    » out '$foo'\n    $foo\n    \n    » out ($foo)\n    bar\n\n## Synonyms\n\n* `export`\n* `!export`\n* `unset`\n\n\n## See Also\n\n* [user-guide/Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by _murex_\n* [user-guide/Variable and Config Scoping](../user-guide/scoping.md):\n  How scoping works within _murex_\n* [commands/`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [commands/`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function (deprecated)\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value (deprecated)"

}
