# `murex-update-exe-list` - Command Reference

> Forces Murex to rescan $PATH looking for exectables

## Description

On application lauch, Murex scans and caches all the executables found in
$PATH on your host. Murex then does regular scans there after. However if
you want to force a new scan (for example you've just installed a new
program and you want it to appear in tab completion) then you can run `murex-update-exe-list`.

## Usage

    murex-update-exe-list

## Examples

    » murex-update-exe-list

## See Also

* [`cpuarch`](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [`cpucount`](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [`os`](../commands/os.md):
  Output the auto-detected OS name