package docs

func init() {

	Definition["(murex named pipe)"] = "# _murex_ Shell Docs\n\n## Command Reference: `<>` (murex named pipe)\n\n> Reads from a _murex_ named pipe\n\n## Description\n\nSometimes you will need to start a command line with a _murex_ named pipe, eg\n\n    » <namedpipe> -> match: foobar\n    \n> See the documentation on `pipe` for more details about _murex_ named pipes.\n\n## Usage\n\nRead from pipe\n\n    <namedpipe> -> <stdout>\n    \nWrite to pipe\n\n    <stdin> -> <namedpipe>\n\n## Examples\n\nThe follow two examples function the same\n\n    » pipe: example\n    » bg { <example> -> match: 2 }\n    » a: <example> [1..3]\n    2\n    » !pipe: example\n\n## Detail\n\n### What are _murex_ named pipes?\n\nIn POSIX, there is a concept of STDIN, STDOUT and STDERR, these are FIFO files\nwhile are \"piped\" from one executable to another. ie STDOUT for application 'A'\nwould be the same file as STDIN for application 'B' when A is piped to B:\n`A | B`. _murex_ adds a another layer around this to enable support for passing\ndata types and builtins which are agnostic to the data serialization format\ntraversing the pipeline. While this does add overhead the advantage is this new\nwrapper can be used as a primitive for channelling any data from one point to\nanother.\n\n_murex_ named pipes are where these pipes are created in a global store,\ndecoupled from any executing functions, named and can then be used to pass\ndata along asynchronously.\n\nFor example\n\n    pipe: example\n    \n    bg {\n        <example> -> match: Hello\n    }\n    \n    out: \"foobar\"        -> <example>\n    out: \"Hello, world!\" -> <example>\n    out: \"foobar\"        -> <example>\n    \n    !pipe: example\n    \nThis returns `Hello, world!` because `out` is writing to the **example** named\npipe and `match` is also reading from it in the background (`bg`).\n\nNamed pipes can also be inlined into the command parameters with `<>` tags\n\n    pipe: example\n    \n    bg {\n        <example> -> match: Hello\n    }\n    \n    out: <example> \"foobar\"\n    out: <example> \"Hello, world!\"\n    out: <example> \"foobar\"\n    \n    !pipe: example\n    \n> Please note this is also how `test` works.\n\n_murex_ named pipes can also represent network sockets, files on a disk or any\nother read and/or write endpoint. Custom builtins can also be written in Golang\nto support different abstractions so your _murex_ code can work with those read\nor write endpoints transparently.\n\nTo see the different supported types run\n\n    runtime --pipes\n    \n### Namespacing and used in modules and packages\n\nPipes created via `pipe` are created in the global namespace. This allows pipes\nto be used across different functions easily however it does pose a risk with\nname clashes where _murex_ named pipes are used heavily. Thus is it recommended\nthat pipes created in modules should be prefixed with the name of its package.\n\n## Synonyms\n\n* `(murex named pipe)`\n* `<>`\n\n\n## See Also\n\n* [commands/`<stdin>` ](../commands/stdin.md):\n  Read the STDIN belonging to the parent code block\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`bg`](../commands/bg.md):\n  Run processes in the background\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`pipe`](../commands/pipe.md):\n  Manage _murex_ named pipes\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [parser/namedpipe](../parser/namedpipe.md):\n  "

}
