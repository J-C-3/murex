package docs

func init() {

	Definition["mtac"] = "# _murex_ Shell Docs\n\n## Command Reference: `mtac`\n\n> Reverse the order of an array\n\n## Description\n\n`mtac` takes input from STDIN and reverses the order of it.\n\nIt's name is derived from a program called `tac` - a tool that functions\nlike `cat` but returns the contents in the reverse order. The difference\nwith the `mtac` builtin is that it is data-type aware. So it doesn't just\nfunction as a replacement for `tac` but it also works on JSON arrays,\ns-expressions, and any other data-type supporting arrays compiled into\n_murex_. \n\n## Usage\n\n    <stdin> -> mtac -> <stdout>\n\n## Examples\n\n    » ja: [Monday..Friday] -> mtac\n    [\n        \"Friday\",\n        \"Thursday\",\n        \"Wednesday\",\n        \"Tuesday\",\n        \"Monday\"\n    ]\n    \n    # Normal output (without mtac)\n    » ja: [Monday..Friday]\n    [\n        \"Monday\",\n        \"Tuesday\",\n        \"Wednesday\",\n        \"Thursday\",\n        \"Friday\"\n    ]\n\n## Detail\n\nPlease bare in mind that while _murex_ is optimised with concurrency and\nstreaming in mind, it's impossible to reverse an incomplete array. Thus all\nall of STDIN must have been read and that file closed before `mtac` can\noutput.\n\nIn practical terms you shouldn't notice any difference except for when\nSTDIN is a long running process or non-standard stream (eg network pipe).\n\n## See Also\n\n* [commands/`2darray` ](../commands/2darray.md):\n  Create a 2D JSON array from multiple input sources\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [commands/`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`prefix`](../commands/prefix.md):\n  Prefix a string to every item in a list\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [commands/`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [commands/`suffix`](../commands/suffix.md):\n  Prefix a string to every item in a list\n* [commands/`ta` (mkarray)](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type\n* [commands/len](../commands/len.md):\n  "

}
