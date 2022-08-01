package docs

func init() {

	Definition["rx"] = "# _murex_ Shell Docs\n\n## Command Reference: `rx`\n\n> Regexp pattern matching for file system objects (eg `.*\\\\.txt`)\n\n## Description\n\nReturns a list of files and directories that match a regexp pattern.\n\nOutput is a JSON list.\n\n## Usage\n\n    rx: pattern -> <stdout>\n    \n    !rx: pattern -> <stdout>\n    \n    <stdin> -> rx: pattern -> \n\n## Examples\n\nInline regex file matching:\n\n    cat: @{ rx: '.*\\.txt' }\n    \nWriting a list of files to disk:\n\n    rx: '.*\\.go' -> > filelist.txt\n    \nChecking if files exist:\n\n    if { rx: somefiles.* } then {\n        # files exist\n    }\n    \nChecking if files do not exist:\n\n    !if { rx: somefiles.* } then {\n        # files do not exist\n    }\n    \nReturn all files apart from text files:\n\n    !g: '\\.txt$'\n\n## Detail\n\n### Traversing Directories\n\nUnlike globbing (`g`) which can traverse directories (eg `g: /path/*`), `rx` is\nonly designed to match file system objects in the current working directory.\n\n`rx` uses Go (lang)'s standard regexp engine.\n\n### Inverse Matches\n\nIf you want to exclude any matches based on wildcards, rather than include\nthem, then you can use the bang prefix. eg\n\n    » rx: READ*                                                                                                                                                              \n    [\n        \"README.md\"\n    ]\n    \n    murex-dev» !rx: .*\n    Error in `!rx` (1,1): No data returned.\n\n## Synonyms\n\n* `rx`\n* `!rx`\n\n\n## See Also\n\n* [commands/`f`](../commands/f.md):\n  Lists or filters file system objects (eg files)\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [commands/`match`](../commands/match.md):\n  Match an exact value in an array\n* [commands/`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings"

}
