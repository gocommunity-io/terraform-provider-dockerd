//go:build windows

package provider

import (
	"os/exec"
	"syscall"
)

func _56af3332aec8(_3325dea0fabf *exec.Cmd) {
	_3325dea0fabf.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x00000200 | 0x08000000,
	}
}
