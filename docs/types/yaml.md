# _murex_ Shell Docs

## Data-Type Reference: `yaml` 

> YAML Ain't Markup Language (YAML)

## Description

YAML support within _murex_ is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of _murex_, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.

## Default Associations

* **Extension**: `yaml`
* **Extension**: `yml`
* **MIME**: `application/x-yaml`
* **MIME**: `application/yaml`
* **MIME**: `text/x-yaml`
* **MIME**: `text/yaml`


## Supported Hooks

* `Marshal()`
    Supported
* `ReadArray()`
    Works with YAML arrays. Maps are converted into arrays
* `ReadIndex()`
    Works against all properties in YAML
* `ReadMap()`
    Works with YAML maps
* `ReadNotIndex()`
    Works against all properties in YAML
* `Unmarshal()`
    Supported
* `WriteArray()`
    Works with YAML arrays

## See Also

* [apis/ReadIndex](../apis/ReadIndex.md):
  
* [apis/ReadNotIndex](../apis/ReadNotIndex.md):
  
* [apis/`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [apis/`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [apis/`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [apis/`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [apis/`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [types/`json` ](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [types/`jsonl` ](../types/jsonl.md):
  JSON Lines (primitive)
* [commands/`open`](../commands/open.md):
  Open a file with a preferred handler
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_