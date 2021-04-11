# _murex_ Shell Docs

## User Guide: _murex_ Profile Files

> A breakdown of the different files loaded on start up

_murex_ has several profile files which are loaded in the following order of
execution:

1. `~/.murex_preload`
2. `~/.murex_modules/*/`
3. `~/.murex_profile`

### `.murex_preload`

This file should only used to define any environmental variables that might
need to be set before the modules are loaded (eg including directories in
`$PATH` if you have anything installed in non-standard locations).

Most of the time this file will be empty bar the standard warning message:

    # This file is loaded before any murex modules. It should only contain
    # environmental variables required for the modules to work eg:
    #
    #     export PATH=...
    #
    # Any other profile config belongs in your profile script instead:
    # /home/$USER/.murex_profile

This file is created upon the first run of _murex_.

### `.murex_modules/`

_murex_'s module directory - where all the modules are installed
to. This directory is managed by `murex-package` builtin.

### `.murex_profile`

This file is comparable to `.bash_profile`, `.bashrc` and `.zshrc` etc. It
is the standard place to put all user and/or machine specific config in.

`.murex_profile` is only read from the users home directory. Unlike bash et
al, profiles will not be read from `/etc/profile.d` nor similar. Modules
should be used in its place.

## See Also

* [user-guide/Modules and Packages](../user-guide/modules.md):
  An introduction to _murex_ modules and packages
* [commands/`murex-package`](../commands/murex-package.md):
  _murex_'s package manager