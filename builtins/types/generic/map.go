package generic

import (
	"bufio"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		row := rxWhitespace.Split(scanner.Text(), -1)
		debug.Json("row", row)
		for i := range row {
			callback(strconv.Itoa(i), row[i], i+1 == len(row))
		}
	}

	err := scanner.Err()
	return err

}
