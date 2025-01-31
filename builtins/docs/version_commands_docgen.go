package docs

func init() {

	Definition["version"] = "# `version`  - Command Reference\n\n> Get Murex version\n\n## Description\n\nReturns Murex version number\n\n## Usage\n\n    version [ flags ] -> <stdout>\n\n## Examples\n\nRan without any parameters\n\n    » version\n    murex: 0.51.1200 BETA\n    \nRan with the `--no-app-name` parameter\n\n    » version --no-app-name\n    0.51.1200 BETA\n    \nRan with the `--short` parameter\n\n    » version --short\n    0.51\n\n## Flags\n\n* `--no-app-name`\n    Returns full version string minus app name\n* `--short`\n    Returns only the major and minor version as a `num` data-type\n\n## See Also\n\n* [`args` ](../commands/args.md):\n  Command line flag parser for Murex shell scripting\n* [`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [`config`](../commands/config.md):\n  Query or define Murex runtime settings\n* [`function`](../commands/function.md):\n  Define a function block\n* [`murex-parser` ](../commands/murex-parser.md):\n  Runs the Murex parser against a block of code \n* [`private`](../commands/private.md):\n  Define a private function block\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of Murex\n* [`source` ](../commands/source.md):\n  Import Murex code from another file of code block"

}
