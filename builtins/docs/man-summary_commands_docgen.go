package docs

func init() {

	Definition["man-summary"] = "# `man-summary` - Command Reference\n\n> Outputs a man page summary of a command\n\n## Description\n\n`man-summary` reads the man pages for a given command and outputs it's\nsummary (if one exists).\n\n## Usage\n\n    man-summary command -> <stdout>\n\n## Examples\n\n    » man-summary: man \n    man - an interface to the on-line reference manuals\n\n## See Also\n\n* [`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [`murex-docs`](../commands/murex-docs.md):\n  Displays the man pages for _murex_ builtins\n* [`summary` ](../commands/summary.md):\n  Defines a summary help text for a command"

}
