#!/usr/bin/env bats -t

@test "List descriptors" {
	run ./bin/psgo -list
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "args, capamb, capbnd, capeff, capinh, capprm, comm, etime, group, groups, hgroup, hgroups, hpid, huid, huser, label, nice, pcpu, pgid, pid, ppid, rgroup, rss, ruser, seccomp, state, stime, time, tty, uid, user, vsz" ]]
}
