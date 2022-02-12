package docs

func init() {

	Definition["["] = "# _murex_ Shell Docs\n\n## Command Reference: `[` (index)\n\n> Outputs an element from an array, map or table\n\n## Description\n\nOutputs an element or multiple elements from an array, map or table.\n\nPlease note that indexes in _murex_ are counted from zero.\n\n## Usage\n\n    <stdin> -> [ element ] -> <stdout>\n    $variable[ element ] -> <stdout>\n    \n    <stdin> -> ![ element ] -> <stdout>\n\n## Examples\n\nReturn the 2nd (1), 4th (3) and 6th (5) element in an array\n\n    » ja [0..9] -> [ 1 3 5 ]\n    [\n        \"1\",\n        \"3\",\n        \"5\"\n    ]\n    \nReturn the data-type and description of **config shell syntax-highlighting**\n\n    » config -> [[ /shell/syntax-highlighting ]] -> [ Data-Type Description ]\n    [\n        \"bool\",\n        \"Syntax highlighting of murex code when in the interactive shell\"\n    ]\n    \nReturn all elements _except_ for 1 (2nd), 3 (4th) and 5 (6th)\n\n    » a: [0..9]-> ![ 1 3 5 ]\n    0\n    2\n    4\n    6\n    7\n    8\n    9\n    \nReturn all elements except for the data-type and description\n\n    » config -> [[ /shell/syntax-highlighting ]] -> ![ Data-Type Description ]\n    {\n        \"Default\": true,\n        \"Dynamic\": false,\n        \"Global\": true,\n        \"Value\": true\n    }\n\n## Detail\n\n### Index counts from zero\n\nIndexes in _murex_ behave like any other computer array in that all arrays\nstart from zero (`0`).\n\n### Include vs exclude\n\nAs demonstrated in the examples above, `[` specifies elements to include\nwhere as `![` specifies elements to exclude.\n\n### Don't error upon missing elements\n\nBy default, **index** generates an error if an element doesn't exist. However\nyou can disable this behavior in `config`\n\n    » config -> [ foobar ]\n    Error in `[` ((builtin) 2,11): Key 'foobar' not found\n    \n    » config set index silent true\n    \n    » config -> [ foobar ]\n\n## Synonyms\n\n* `[`\n* `![`\n\n\n## See Also\n\n* [commands/`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [commands/len](../commands/len.md):\n  "

}
