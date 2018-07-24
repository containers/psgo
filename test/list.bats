#!/usr/bin/env bats -t

@test "List descriptors" {
	run ./bin/psgo -list
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "args, capbnd, capeff, capinh, capprm, comm, etime, group, label, nice, pcpu, pgid, pid, ppid, rgroup, ruser, seccomp, time, tty, user, vsz" ]]
}
