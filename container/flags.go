package container

import "syscall"

const (
	CLONE_FLAGS = syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC |
		syscall.CLONE_NEWNET | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS
)
