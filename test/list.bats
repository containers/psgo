#!/usr/bin/env bats -t

@test "List descriptors" {
	run ./bin/psgo -list
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "args, capamb, capbnd, capeff, capinh, capprm, comm, etime, group, hgroup, hpid, huser, label, nice, pcpu, pgid, pid, ppid, rgroup, ruser, seccomp, state, time, tty, user, vsz" ]]
}
