package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_all.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_all.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "function h {\n    # Output the murex history log in a human readable format\n    history -> foreach { -> set json line; out \"$line[Index]: $line[Block]\" } -> cast *\n}\n\nfunction aliases {\n\t# Output the aliases in human readable format\n\truntime: --aliases -> formap name alias {\n        $name -> sprintf: \"%10s => ${esccli @alias}\\n\"\n\t} -> cast: str\n}\n\ntest unit function aliases {\n    \"PreBlock\": ({\n        alias ALIAS_UNIT_TEST=example param1 param2 param3\n    }),\n    \"StdoutRegex\": \"([- _0-9a-zA-Z]+ => .*?\\n)+\",\n    \"StdoutType\": \"str\",\n    \"PostBlock\": ({\n        !alias ALIAS_UNIT_TEST\n    })\n}\n\nautocomplete: set cd { [{\n    \"IncDirs\": true\n}] }\n\nautocomplete: set mkdir { [{\n    \"IncDirs\": true,\n    \"AllowMultiple\": true\n}] }\n\nautocomplete: set rmdir { [{\n    \"IncDirs\": true,\n    \"AllowMultiple\": true\n}] }\n\nautocomplete: set exec { [\n    {\n        \"IncFiles\": true,\n        \"IncDirs\": true,\n        \"IncExePath\": true\n    },\n    {\n        \"NestedCommand\": true\n    }\n] }\n\nautocomplete: set format { [{\n    \"Dynamic\": ({ runtime: --marshallers })\n}] }\n\nautocomplete: set swivel-datatype { [{\n    \"Dynamic\": ({ runtime: --marshallers })\n}] }\n\nprivate autocomplete.data-types {\n    # Returns all murex data-types compiled\n    runtime: --readarray -> format: str\n    runtime: --writearray -> format: str\n    runtime: --readmap -> format: str\n    runtime: --marshallers -> format: str\n    runtime: --unmarshallers -> format: str\n}\n\ntest unit private autocomplete.data-types {\n    \"StdoutRegex\": (^(([a-z0-9]+|\\*)\\n)+),\n    \"StdoutType\":  \"str\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsArray\": true\n}\n\nautocomplete: set cast { [{\n    \"Dynamic\": ({ autocomplete.data-types })\n}] }\n\nautocomplete: set tout { [{\n    \"Dynamic\": ({ autocomplete.data-types })\n}] }\n\nprivate autocomplete.variables.locals {\n    # Returns all local variable names\n    runtime: --variables -> formap k v { out $k } -> cast: str\n}\n\nprivate autocomplete.variables.globals {\n    # Returns all global variable names\n    runtime: --globals -> formap k v { out $k } -> cast: str\n}\n\nprivate autocomplete.variables {\n    # Returns all global and local variable names\n    autocomplete.variables.locals\n    autocomplete.variables.globals\n}\n\ntest unit private autocomplete.variables {\n    \"PreBlock\": ({ global MUREX_UNIT_TEST=foobar }),\n    \"PostBlock\": ({ !global MUREX_UNIT_TEST }),\n    \"StdoutRegex\": (^([_a-zA-Z0-9]+\\n)+),\n    \"StdoutType\":  \"str\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsArray\": true\n}\n\nautocomplete: set set { [{\n    \"Dynamic\": ({ autocomplete.variables.locals })\n}] }\n\nautocomplete: set !set { [{\n    \"Dynamic\": ({ autocomplete.variables.locals })\n}] }\n\nautocomplete: set global { [{\n    \"Dynamic\": ({ autocomplete.variables.globals })\n}] }\n\nautocomplete: set !global { [{\n    \"Dynamic\": ({ autocomplete.variables.globals })\n}] }\n\nautocomplete: set export { [{\n    \"Dynamic\": ({ autocomplete.variables })\n}] }\n\nautocomplete: set !export { [{\n    \"DynamicDesc\": ({ runtime --exports })\n}] }\n\nautocomplete: set unset { [{\n    \"DynamicDesc\": ({ runtime --exports })\n}] }\n\nautocomplete: set \"[\" { [{\n    \"AnyValue\": true,\n    \"AllowMultiple\": true,\n    \"ExecCmdline\": true,\n    \"Dynamic\": ({\n        switch ${ get-type: stdin } {\n            case * {\n                #<stdin> -> tabulate -> [ 0: ] -> format json -> [ 0 ]\n                <stdin> -> [ 0: ] -> format json -> [ 0 ] -> append \"]\"\n            }\n\n            case csv {\n                <stdin> -> [ 0: ] -> format json -> [ 0 ] -> append \"]\"\n            }\n            \n            case jsonl {\n                <stdin> -> [ 0 ] -> set header\n                $header -> cast utf8 -> [ 0 -1 ] -> set jsonl_format\n                if { = jsonl_format==`[]` } then {\n                    tout json $header -> append \"]\"\n                }\n            }\n\n            catch {\n                <stdin> -> formap k v { out $k } -> cast str -> append \"]\"\n            }\n        }\n    })\n}] }\n\nautocomplete: set \"[[\" { [{\n    \"AnyValue\": true,\n    \"ExecCmdline\": true,\n    \"AutoBranch\": true,\n    \"Dynamic\": ({ -> struct-keys -> append \"]]\" })\n} ]}\n\nprivate autocomplete.config.get.apps {\n    # Returns all app names for the 'app' field in `config`\n    config: -> formap k v { out $k } -> cast: str -> msort\n}\n\ntest unit private autocomplete.config.get.apps {\n    \"StdoutRegex\": (shell),\n    \"StdoutType\":  \"str\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsArray\": true\n}\n\nprivate autocomplete.config.get.keys {\n    # Returns all keys for the 'app' field in `config`\n    config -> [ $ARGS[1] ] -> formap k v { out $k } -> cast: str -> msort\n}\n\ntest unit private autocomplete.config.get.keys {\n    \"Parameters\": [ \"shell\" ],\n    \"StdoutRegex\": (prompt),\n    \"StdoutType\":  \"str\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsArray\": true\n}\n\nautocomplete: set config { [{\n    \"Flags\": [ \"get\", \"set\", \"eval\", \"define\", \"default\" ],\n    \"FlagValues\": {\n        \"get\": [\n            { \"Dynamic\": ({ autocomplete.config.get.apps }) },\n            { \"Dynamic\": ({ autocomplete.config.get.keys $ARGS[2] }) }\n        ],               \n        \"set\": [\n            { \"Dynamic\": ({ autocomplete.config.get.apps }) },\n            { \"Dynamic\": ({ autocomplete.config.get.keys $ARGS[2] }) },\n            { \"Dynamic\": ({\n\t\t\t\tswitch {\n\t\t\t\t\tcase { = `${ config -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ Data-Type ]}`==`bool` } {\n\t\t\t\t\t\tja [true,false]\n\t\t\t\t\t}\n\n\t\t\t\t\tcase { config -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ <!null> Options ] } {\n\t\t\t\t\t\tconfig -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ Options ]\n\t\t\t\t\t}\n\t\t\t\t\t\n                \tcatch {\n\t\t\t\t\t\tout ${ config -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ Default ]}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}) }\n        ],\n        \"eval\": [\n            { \"Dynamic\": ({ autocomplete.config.get.apps }) },\n            { \"Dynamic\": ({ autocomplete.config.get.keys $ARGS[2] }) }\n        ],\n        \"default\": [\n            { \"Dynamic\": ({ autocomplete.config.get.apps }) },\n            { \"Dynamic\": ({ autocomplete.config.get.keys $ARGS[2] }) }\n        ]\n    }\n}] }\n\nautocomplete: set !config { [\n    { \"Dynamic\": ({ autocomplete.config.get.apps }) },\n    { \"Dynamic\": ({ autocomplete.config.get.keys $ARGS[1] }) }\n] }\n\n\nautocomplete: set event { [\n    {\n        \"Dynamic\": ({ runtime: --events -> formap k v { out $k } })\n    }\n] }\n\nautocomplete: set !event { [\n    {\n        \"Dynamic\": ({ runtime: --events -> formap k v { out $k } -> msort })\n    },\n    {\n        \"Dynamic\": ({ runtime: --events -> [ $ARGS[1] ] -> formap k v { out $k } -> msort })\n    }\n] }\n\nprivate autocomplete.alias {\n    # Returns a map of all alises and the code they execute\n    runtime: --aliases -> formap --jmap name value { $name } { out @value }\n}\n\ntest unit private autocomplete.alias {\n    \"StdoutRegex\": (jobs),\n\t\"StdoutType\":  \"json\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsMap\": true\n}\n\nautocomplete: set !alias { [{\n    \"DynamicDesc\": ({ autocomplete.alias }),\n    \"ListView\": true\n}] }\n\nprivate autocomplete.functions {\n    # Returns a map of all murex public functions\n    runtime: --functions -> formap --jmap k v { $k } { out: $v[summary] }\n}\n\ntest unit private autocomplete.functions {\n    \"PreBlock\": ({\n        function unit.test.autocomplete.functions {\n            out \"This is only a dummy function for testing\"\n        }\n    }),\n    \"PostBlock\": ({\n        !function unit.test.autocomplete.functions\n    }),\n    \"StdoutRegex\": (unit.test.autocomplete.functions),\n\t\"StdoutType\":  \"json\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsMap\": true\n}\n\nautocomplete: set !function { [{\n    \"DynamicDesc\": ({ autocomplete.functions }),\n    \"ListView\": true\n} ]}\n\nprivate autocomplete.privates {\n    # Returns a map of all murex private functions\n    runtime: --privates  -> struct-keys: 3 -> regexp m,/.*?/.*?/, -> foreach --jmap private { $private } { runtime: --privates -> [[ $private/Summary ]] }\n}\n\ntest unit private autocomplete.privates {\n\t\"StdoutType\":  \"json\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsMap\": true\n}\n\nprivate autocomplete.builtins {\n    # Returns a map of all murex builtins\n    runtime --builtins -> foreach --jmap builtin { $builtin } { murex-docs --summary $builtin }\n}\n\ntest unit private autocomplete.builtins {\n\t\"StdoutType\":  \"json\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsMap\": true\n}\n\nautocomplete: set autocomplete { [{\n    \"Flags\": [ \"get\", \"set\" ],\n    \"FlagValues\": {\n        \"get\": [{\n            \"Dynamic\": ({\n                runtime: --autocomplete -> formap: cmd ! { out $cmd } -> cast: str\n            })\n        }]\n    }\n}] }\n\nprivate git-branch {\n    # Returns a list of branches excluding currect checked-out branch in the current git repository\n    git branch -> [ :0 ] -> !match *\n}\n\nautocomplete: set git { [{\n    #\"Flags\": [ \"clone\", \"init\", \"add\", \"mv\", \"reset\", \"rm\", \"bisect\", \"grep\", \"log\", \"show\", \"status\", \"branch\", \"checkout\", \"commit\", \"diff\", \"merge\", \"rebase\", \"tag\", \"fetch\", \"pull\", \"push\", \"stash\" ],\n    \"DynamicDesc\": ({\n        git: help -a -> @[..^Ancillary]re -> tabulate: --map\n    }),\n    \"ListView\": true,\n    \"FlagValues\": {\n        \"init\": [{\n            \"Flags\": [\"--bare\"]\n        }],\n        \"add\": [{\n            #\"IncFiles\": true,\n            \"AllowMultiple\": true,\n            \"Dynamic\": ({\n                git status -s -> regexp 'f/^.[^\\s] [\"]?(.*?)[\"]?$/' -> cast str\n            })\n        }],\n        \"diff\": [{\n            #\"IncFiles\": true,\n            \"AllowMultiple\": true,\n            \"Dynamic\": ({\n                git status -s -> [:1]\n            })\n        }],\n        \"mv\": [{ \n            \"IncFiles\": true\n        }],\n        \"rm\": [{\n            \"IncFiles\": true,\n            \"AllowMultiple\": true\n        }],\n        \"checkout\": [{\n            \"Dynamic\": ({ git-branch }),\n            \"Flags\": [ \"-b\" ]\n        }],\n        \"merge\": [{\n            \"Dynamic\": ({ git-branch })\n        }],\n        \"commit\": [{\n            \"Flags\": [\"-a\", \"-m\", \"--amend\"],\n            \"FlagValues\": {\n                \"--amend\": [{ \"AnyValue\": true }]\n            },\n            \"AllowMultiple\": true\n        }]\n    }\n}] }\n\nautocomplete: set docker { [\n    {\n        \"DynamicDesc\": ({\n            docker help -> @[^Usage:..]re -> tabulate: --split-comma --map\n        }),\n\n        #\"AllowMultiple\": true,\n        #\"AnyValue\": true,\n        \"ListView\": true,\n\n        \"FlagValues\": {\n            \"builder\": [{\n                \"DynamicDesc\": ({\n                    docker help builder -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"config\": [{\n                \"DynamicDesc\": ({\n                    docker help config -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"container\": [{\n                \"DynamicDesc\": ({\n                    docker help container -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"context\": [{\n                \"DynamicDesc\": ({\n                    docker help context -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"image\": [{\n                \"DynamicDesc\": ({\n                    docker help image -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"network\": [{\n                \"DynamicDesc\": ({\n                    docker help network -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"node\": [{\n                \"DynamicDesc\": ({\n                    docker help node -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"plugin\": [{\n                \"DynamicDesc\": ({\n                    docker help plugin -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"secret\": [{\n                \"DynamicDesc\": ({\n                    docker help secret -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"service\": [{\n                \"DynamicDesc\": ({\n                    docker help service -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"stack\": [{\n                \"DynamicDesc\": ({\n                    docker help stack -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"swarm\": [{\n                \"DynamicDesc\": ({\n                    docker help swarm -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"system\": [{\n                \"DynamicDesc\": ({\n                    docker help system -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"trust\": [{\n                \"DynamicDesc\": ({\n                    docker help trust -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }],\n\n            \"volume\": [{\n                \"DynamicDesc\": ({\n                    docker help volume -> @[^Usage:..]re -> tabulate: --split-comma --map\n                })\n            }]\n        }\n    },\n    {\n        \"IncFiles\": true\n    }\n] }\n\nprivate autocomplete.docker-compose.services {\n    # Returns a list of services described in docker-compose.yaml\n    open docker-compose.yaml -> [ services ] -> formap k v { out \"$k\" } -> cast str\n}\n\nautocomplete: set docker-compose { [{\n    \"Flags\": [\"build\",\"bundle\",\"config\",\"create\",\"down\",\"events\",\"exec\",\"help\",\"images\",\"kill\",\"logs\",\"pause\",\"port\",\"ps\",\"pull\",\"push\",\"restart\",\"rm\",\"run\",\"scale\",\"start\",\"stop\",\"top\",\"unpause\",\"up\",\"version\"],\n    \"FlagValues\": {\n        \"build\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"create\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"events\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"exec\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"kill\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"logs\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"pause\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"pull\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"push\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"restart\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"run\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"scale\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"start\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"stop\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"top\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"unpause\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }],\n        \"up\": [{\n            \"Dynamic\": ({ autocomplete.docker-compose.services })\n        }]\n    }\n}] }\n\nautocomplete: set terraform { [{\n    \"Flags\": [\"apply\",\"console\",\"destroy\",\"env\",\"fmt\",\"get\",\"graph\",\"import\",\"init\",\"output\",\"plan\",\"providers\",\"push\",\"refresh\",\"show\",\"taint\",\"untaint\",\"validate\",\"version\",\"workspace\"],\n    \"FlagValues\": {\n        \"workspace\": [\n            {\n                \"Flags\": [ \"new\", \"delete\", \"select\", \"list\", \"show\" ]\n            }\n        ]\n    }\n}] }\n\nautocomplete: set gopass { [{\n    \"Flags\": [\"--yes\",\"--clip\",\"-c\",\"--help\",\"-h\",\"--version\",\"-v\"],\n    \"AllowMultiple\": true,\n    \"Dynamic\": ({ exec: @ARGS --generate-bash-completion }),\n    \"AutoBranch\": true\n}] }\n\nautocomplete: set debug { [{\n    \"Flags\": [\"on\", \"off\"]\n}] }\n\nautocomplete: set murex-package {\n    [{\n        \"FlagsDesc\": {\n            \"install\": \"Installs a package from a user defined URI\",\n            \"update\":  \"Updates all installed packages\",\n            \"import\":  \"Import packages described in a backup package DB from user defined URI or local path\",\n            \"enable\":  \"Enables a previously disabled package or module\",\n            \"disable\": \"Disables a previously enabled package or module\",\n            \"reload\":  \"Reloads all enabled modules\",\n            \"status\":  \"Returns the version status of locally installed packages\",\n            \"list\":    \"Returns a list of indexed packages/modules (eg what's enabled or disabled)\",\n            \"cd\":      \"Changes working directory to a package's install location\"\n        },\n        \"FlagValues\": {\n            \"import\": [{\n                \"IncFiles\": true\n            }],\n            \"enable\": [{\n                \"DynamicDesc\": ({ murex-package: list disabled }),\n                \"ListView\": true,\n                \"AutoBranch\": true\n            }],\n            \"disable\": [{\n                \"DynamicDesc\": ({ murex-package: list enabled }),\n                \"ListView\": true,\n                \"AutoBranch\": true\n            }],\n            \"list\": [{\n                \"Flags\": [ \"enabled\", \"disabled\", \"loaded\", \"not-loaded\", \"packages\" ]\n            }],\n            \"cd\": [{\n                \"Dynamic\": ({ murex-package: list packages })\n            }]\n        }\n    }]\n}\n\nalias: builtins=runtime --builtins\n\nprivate: test.alias.builtins {\n    # Wrapper function around the alias for `builtins` for unit testing\n    builtins\n}\n\ntest: unit private test.alias.builtins {\n    \"StdoutRegex\": (\"[a-z0-9]+\",),\n    \"StdoutType\":  \"json\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsArray\": true\n}\n\nautocomplete: set murex-docs { [{\n    \"Dynamic\": ({ builtins }),\n\t\"Flags\": [ \"--summary\" ],\n\t\"FlagValues\": {\n\t\t\"--summary\": [{\n\t\t\t\"Dynamic\": ({ builtins })\n\t\t}]\n\t}\n}] }\n\nprivate: autocomplete.aliases.and.builtins {\n    # Returns a list of aliases and builtins\n    runtime: --aliases -> formap k ! { out: $k } -> cast str\n    builtins -> format str\n}\n\ntest: unit private autocomplete.aliases.and.builtins {\n    \"StdoutType\":  \"str\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    }),\n    \"StdoutIsArray\": true\n}\n\nautocomplete: set summary { [{\n    \"IncExePath\": true,\n    \"Dynamic\": ({\n        autocomplete.aliases.and.builtins\n    })\n}] }\n\nconfig: eval shell safe-commands {\n    -> alter --merge / ([\n        \"builtins\", \"jobs\"\n    ])\n}\n\n!if { man-summary terraform } then {\n    summary terraform \"Write, Plan, and Create Infrastructure as Code\"\n}\n\nautocomplete: set terraform {[\n    {\n        \"DynamicDesc\": ({\n            terraform -help @{ $ARGS -> @[1..] } -> tabulate: --map --column-wraps --key-inc-hint\n        }),\n        \"AllowMultiple\": true,\n        \"ListView\": true\n    }\n]}\n\n!if { man-summary go } then {\n    summary go \"Go is a tool for managing Go source code\"\n}\n\n!if { man-summary atom } then {\n    summary atom \"Github Atom - Text editor / IDE\"\n}\n\n!if { man-summary code } then {\n    summary code \"Microsoft Visual Studio Code - Text editor / IDE\"\n}\n\nautocomplete: set zfs {\n    [{\n        \"Dynamic\": ({\n             zfs ? egrep \"^\\t[a-z]+\" -> regexp 'f/\\t+([a-z]+)/' -> uniq \n        })\n    }]\n}\n\nautocomplete: set zpool {\n    [{\n        \"Dynamic\": ({\n             zpool ? egrep \"^\\t[a-z]+\" -> regexp 'f/\\t+([a-z]+)/' -> uniq \n        })\n    }]\n}\n\nfunction sprintf {\n    # This is a wrapper around the OS's default `printf`, replacing the now deprecated builtin of the same name\n    -> set params\n    switch {\n        case  { $params -> ! } { err: \"No parameters passed via STDIN\" }\n        case  { $ARGS[1]-> ! } { err: \"No string supplied to print\"    }\n        catch {\n            printf \"$ARGS[1]\\n\" @params\n        }\n    }\n}\n\nautocomplete set docgen { [\n    {\n        \"AllowMultiple\": true,\n        \"Optional\": true,\n        \"FlagsDesc\": {\n            \"-panic\": \"Write a stack trace on error\",\n            \"-readonly\": \"Don't write output to disk. Use this to test the config\",\n            \"-verbose\": \"Verbose output (all log messages inc warnings)\",\n            \"-version\": \"Output docgen version number and exit\",\n            \"-warning\": \"Display warning messages (will also return a non-zero exit status if warnings found)\"\n        }\n    },\n    {\n        \"FlagsDesc\": {\n            \"-config\": \"Location of the base docgen config file\"\n        },\n        \"FlagValues\": {\n            \"-config\": [{\n                \"IncFiles\": true\n            }]\n        }\n    }\n] }\n\nconfig define open image {\n    \"Description\":  \"Which mode to render images to the terminal.\",\n    \"DataType\":     \"str\",\n    \"Default\":      \"auto\",\n    \"Options\":      [ \"auto\", \"compatible\", \"kitty\", \"iterm\", \"terminology\", \"sixel\" ]\n}\n\nopenagent set image {\n    $ARGS[1] -> set file\n    config get open image -> set mode\n\n    if { = mode==`auto` } {\n        switch {\n            case { $KITTY_WINDOW_ID }      { set mode=kitty }\n            case { = TERM==`xterm-kitty` } { set mode=kitty }\n            case { $TERMINOLOGY }          { set mode=terminology }\n            #case { = TERM_PROGRAM==`iTerm.app` } { set mode=iterm }\n            #case { $ITERM_PROFILE }        { set mode=iterm }\n            catch                          { set mode=compatible }\n        }\n    }\n\n    # If Kitty but running inside a screen / tmux session, fallback to compatible.\n    # We do this because tmux doesn't support reporting screen sizes via the TIOCGWINSZ ioctl.\n    if { and { = mode==`kitty` } { = TERM==`screen` } } {\n        set mode=compatible\n    }\n\n    switch $mode {\n        case compatible {\n            open-image $file\n        }\n\n        case kitty {\n            try {\n                kitty icat $file\n            }\n            catch {\n                err \"Error running `kitty`. Please check `kitty` is installed.\"\n            }\n        }\n\n        case iterm {\n            #out \"{ESC}]1337;File=$file{^G}\"\n            out \"$file\"\n        }\n\n        case terminology {\n            try {\n                tycat -c $file\n            }\n            catch {\n                err \"Error running `tycat`. Please check `terminology` is installed.\"\n            }\n        }\n\n        case sixel {\n            try {\n                img2sixel $file\n            }\n            catch {\n                err \"Error running `img2sixel`. Please check `libsixel` is installed.\"\n            }\n        }\n\n        catch {\n            out \"Invalid rendering method. Please define in `config set open image` - 'auto' is recommended.\"\n        }\n    }\n}\n\nif { microk8s } then {\n\n    private autocomplete.microk8s {\n        # Top level completion results for microk8s\n        microk8s -> tabulate -> cast str\n    }\n\n    test: unit private autocomplete.microk8s {\n        \"StdoutType\":  \"str\",\n        \"StdoutBlock\": ({\n            -> len -> set len;\n            if { = len>0 } then {\n                out \"Len greater than 0\"\n            } else {\n                err \"No elements returned\"\n            }\n        }),\n        \"StdoutIsArray\": true\n    }\n\n    private autocomplete.microk8s.kubectl {\n        # Top level completion results for microk8s kubectl\n        microk8s kubectl help -> tabulate: --map\n    }\n\n    test: unit private autocomplete.microk8s.kubectl {\n        \"StdoutType\":  \"json\",\n        \"StdoutBlock\": ({\n            -> len -> set len;\n            if { = len>0 } then {\n                out \"Len greater than 0\"\n            } else {\n                err \"No elements returned\"\n            }\n        }),\n        \"StdoutIsMap\": true\n    }\n\n    private autocomplete.microk8s.kubectl.dyndesc {\n        # Dynamic completion for microk8s kubectl\n        microk8s kubectl help $ARGS[1] -> @[^Options..^Usage]re -> regexp \"s/:/\\t/\" -> tabulate: --key-inc-hint --map --split-comma\n    }\n\n    test: unit private autocomplete.microk8s.kubectl.dyndesc {\n        \"StdoutType\":  \"json\",\n        \"StdoutBlock\": ({\n            -> len -> set len;\n            if { = len>0 } then {\n                out \"Len greater than 0\"\n            } else {\n                err \"No elements returned\"\n            }\n        }),\n        \"StdoutIsMap\": true\n    }\n\n    autocomplete: set microk8s {[\n        {\n            \"Dynamic\": ({ autocomplete.microk8s }),\n            \"AllowMultiple\": false,\n            \"FlagValues\": {\n                \"kubectl\": [\n                    {\n                        \"DynamicDesc\":   ({ autocomplete.microk8s.kubectl })\n                    },\n                    {\n                        \"DynamicDesc\":   ({ autocomplete.microk8s.kubectl.dyndesc $ARGS[2] }),\n                        \"AllowMultiple\": true,\n                        \"AnyValue\":      true\n                    }\n                ]\n            }\n        }\n    ]}\n\n}")
}
