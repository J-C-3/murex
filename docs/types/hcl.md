# `hcl`  - Data-Type Reference

> HashiCorp Configuration Language (HCL)

## Description

The description below is taken from the HCL git repository's [README](https://github.com/hashicorp/hcl):

> HCL (HashiCorp Configuration Language) is a configuration language built by
> HashiCorp. The goal of HCL is to build a structured configuration language
> that is both human and machine friendly for use with command-line tools, but
> specifically targeted towards DevOps tools, servers, etc.
>
> HCL is also fully JSON compatible. That is, JSON can be used as completely
> valid input to a system expecting HCL. This helps makes systems interoperable
> with other systems.

HCL support within Murex is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of Murex, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.

## Examples

    terraform {
      required_version = "~> 0.11.5"
    }
    
    data "aws_availability_zones" "available" {}
    
    data "aws_vpc" "vpc" {
      id = "${var.vpc_id}"
    }
    
    data "aws_route53_zone" "external" {
      zone_id = "${var.external_hosted_zone_id}"
    }
    
    data "aws_iam_policy_document" "assume_role_policy" {
      statement {
        actions = ["sts:AssumeRole"]
        effect  = "Allow"
    
        principals {
          identifiers = ["ec2.amazonaws.com"]
          type        = "Service"
        }
      }
    }
    
See the HashiCorp's [documentation](https://github.com/hashicorp/hcl) for HCL syntax.

## Default Associations

* **Extension**: `hcl`
* **Extension**: `tf`
* **Extension**: `tfvars`
* **MIME**: `application/hcl`
* **MIME**: `application/x-hcl`
* **MIME**: `text/hcl`
* **MIME**: `text/x-hcl`


## Supported Hooks

* `Marshal()`
    Supported via a JSON marshaller because HCL is designed to be written by humans but "compiled" into JSON
* `ReadArray()`
    Works with HCL arrays. Maps are converted into arrays
* `ReadArrayWithType()`
    Works with HCL arrays. Maps are converted into arrays. Elements data-type in Murex mirrors the HCL type of the element
* `ReadIndex()`
    Works against all properties in HCL
* `ReadMap()`
    Works with HCL maps
* `ReadNotIndex()`
    Works against all properties in HCL
* `Unmarshal()`
    Supported
* `WriteArray()`
    Works with HCL arrays

## See Also

* [`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`ReadArray()` (type)](../apis/ReadArray.md):
  Read from a data type one array element at a time
* [`ReadIndex()` (type)](../apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [`ReadMap()` (type)](../apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [`WriteArray()` (type)](../apis/WriteArray.md):
  Write a data type, one array element at a time
* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`json` ](../types/json.md):
  JavaScript Object Notation (JSON)
* [`jsonl` ](../types/jsonl.md):
  JSON Lines
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`yaml` ](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [mxjson](../types/mxjson.md):
  Murex-flavoured JSON (deprecated)