package errtrace

import (
	"fmt"
	"runtime"
)

func AddTrace(err error) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("%w %s:%d", err, file, line)
}
