package reader

import (
	"bufio"
	"os"
)

var StreamReader = bufio.NewReader(os.Stdin)
