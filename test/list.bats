#!/usr/bin/env bats -t

@test "List descriptors" {
	run ./bin/psgo -list
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "pcpu, group, ppid, user, args, comm, rgroup, nice, pid, pgid, etime, ruser, time, tty, vsz, capinh, capprm, capeff, capbnd, seccomp" ]]
}
