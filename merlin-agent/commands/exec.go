//go:build !windows

/*
Merlin is a post-exploitation command and control framework.

This file is part of Merlin.
Copyright (C) 2024 Russel Van Tuyl

Merlin is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
any later version.

Merlin is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Merlin.  If not, see <http://www.gnu.org/licenses/>.
*/

package commands

import (
	// Standard
	"errors"
	"fmt"
	"os/exec"
)

// ExecuteCommand is a function used to instruct an agent to execute a command on the host operating system
func executeCommand(name string, args []string) (stdout string, stderr string) {
	cmd := exec.Command(name, args...) // #nosec G204

	out, err := cmd.CombinedOutput()
	if cmd.Process != nil {
		stdout = fmt.Sprintf("Created %s process with an ID of %d\n", name, cmd.Process.Pid)
	}
	stdout += string(out)

	if err != nil {
		stderr = err.Error()
	}

	return stdout, stderr
}

// ExecuteShellcodeSelf executes provided shellcode in the current process
//
//lint:ignore SA4009 Function needs to mirror exec_windows.go and inputs must be used
func ExecuteShellcodeSelf(shellcode []byte) error {
	return errors.New("shellcode execution is not implemented for this operating system")
}

// ExecuteShellcodeRemote executes provided shellcode in the provided target process
//
//lint:ignore SA4009 Function needs to mirror exec_windows.go and inputs must be used
func ExecuteShellcodeRemote(shellcode []byte, pid uint32) error {
	return errors.New("shellcode execution is not implemented for this operating system")
}

// ExecuteShellcodeRtlCreateUserThread executes provided shellcode in the provided target process using the Windows RtlCreateUserThread call
//
//lint:ignore SA4009 Function needs to mirror exec_windows.go and inputs must be used
func ExecuteShellcodeRtlCreateUserThread(shellcode []byte, pid uint32) error {
	return errors.New("shellcode execution is not implemented for this operating system")
}

// ExecuteShellcodeQueueUserAPC executes provided shellcode in the provided target process using the Windows QueueUserAPC API call
//
//lint:ignore SA4009 Function needs to mirror exec_windows.go and inputs must be used
func ExecuteShellcodeQueueUserAPC([]byte, uint32) error {
	return errors.New("shellcode execution is not implemented for this operating system")
}

// ExecuteShellcodeCreateProcessWithPipe creates a child process, redirects STDOUT/STDERR to an anonymous pipe, injects/executes shellcode, and retrieves output
//
//lint:ignore SA4009 Function needs to mirror exec_windows.go and inputs must be used
func ExecuteShellcodeCreateProcessWithPipe(string, string, string) (stdout string, stderr string, err error) {
	return stdout, stderr, fmt.Errorf("CreateProcess modules in not implemented for this operating  system")
}

// miniDump is a Windows only module function to dump the memory of the provided process
//
//lint:ignore SA4009 Function needs to mirror exec_windows.go and inputs must be used
func miniDump(string, string, uint32) (map[string]interface{}, error) {
	var mini map[string]interface{}
	return mini, errors.New("minidump doesn't work on non-windows hosts")
}
