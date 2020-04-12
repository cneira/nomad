// +build illumos

package term // import "github.com/docker/docker/pkg/term"

import (
	"syscall"
	"unsafe"
)
// #include <termios.h>
// #include <sys/ioctl.h>
// int tcgetattr(int fd, struct termios *);
// int tcgetattr(int fd, struct termios *);
import "C"

func tcget(fd uintptr, p *Termios) syscall.Errno {
if err := C.tcgetattr(C.int(fd), (*C.struct_termios)(unsafe.Pointer(p))); err != 0 {
		return (syscall.Errno)(err)
	}
		return  0
}

func tcset(fd uintptr, p *Termios) syscall.Errno {
if err := C.tcsetattr(C.int(fd),C.TCSANOW, (*C.struct_termios)(unsafe.Pointer(p))); err != 0 {
		return (syscall.Errno)(err)
	}
return 0 
}
