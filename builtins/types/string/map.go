package string

import (
	"bufio"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"strconv"
	"strings"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	scanner := bufio.NewScanner(read)
	i := -1
	for scanner.Scan() {
		i++
		callback(strconv.Itoa(i), strings.TrimSpace(string(scanner.Bytes())), false)
	}

	return scanner.Err()
}
