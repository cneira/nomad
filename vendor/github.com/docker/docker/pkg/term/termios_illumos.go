// +build illumos

package term // import "github.com/docker/docker/pkg/term"

// include <sys/types.h>
// include <termios.h>

import (
	"errors"
)
// #include <termios.h>
// #include <sys/ioctl.h>
// int tcgetattr(int fd, struct termios *);
import "C"

// Termios is the Unix API for terminal I/O.
type Termios C.struct_termios

// MakeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func MakeRaw(fd uintptr) (*State, error) {
	var oldState C.struct_termios 
	var oldState2 State 
	if err := C.tcgetattr(C.int(fd), &oldState); err != 0 {
		return nil, errors.New("err")
	}

	newState := oldState
	newState.c_iflag &^= (C.IGNBRK | C.BRKINT | C.PARMRK | C.ISTRIP | C.INLCR |C.IGNCR | C.ICRNL | C.IXON)
	newState.c_oflag &^= C.OPOST
	newState.c_lflag &^= (C.ECHO | C.ECHONL | C.ICANON | C.ISIG | C.IEXTEN)
	newState.c_cflag &^= (C.CSIZE | C.PARENB)
	newState.c_cflag |= C.CS8
	newState.c_cc[C.VMIN] = 1
	newState.c_cc[C.VTIME] = 0

	return &oldState2, nil
}
