#!/usr/bin/env bats -t

@test "Default header" {
	run ./bin/psgo
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "USER" ]]
	[[ ${lines[0]} =~ "PID" ]]
	[[ ${lines[0]} =~ "PPID" ]]
	[[ ${lines[0]} =~ "%CPU" ]]
	[[ ${lines[0]} =~ "ELAPSED" ]]
	[[ ${lines[0]} =~ "TTY" ]]
	[[ ${lines[0]} =~ "TIME" ]]
	[[ ${lines[0]} =~ "COMMAND" ]]
}

@test "%CPU header" {
	run ./bin/psgo -format "%C"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "%CPU" ]]

	run ./bin/psgo -format "pcpu"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "%CPU" ]]
}

@test "GROUP header" {
	run ./bin/psgo -format "%G"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "GROUP" ]]

	run ./bin/psgo -format "group"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "GROUP" ]]
}

@test "GROUPS header" {
	run ./bin/psgo -format "groups"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "GROUPS" ]]
}

@test "PPID header" {
	run ./bin/psgo -format "%P"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "PPID" ]]

	run ./bin/psgo -format "ppid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "PPID" ]]
}

@test "USER header" {
	run ./bin/psgo -format "%U"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "USER" ]]

	run ./bin/psgo -format "user"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "USER" ]]
}

@test "UID header" {
	run ./bin/psgo -format "uid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "UID" ]]
}

@test "COMMAND (args) header" {
	run ./bin/psgo -format "%a"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "COMMAND" ]]

	run ./bin/psgo -format "args"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "COMMAND" ]]
}

@test "COMMAND (comm) header" {
	run ./bin/psgo -format "%c"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "COMMAND" ]]

	run ./bin/psgo -format "comm"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "COMMAND" ]]
}

@test "RGROUP header" {
	run ./bin/psgo -format "%g"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "RGROUP" ]]

	run ./bin/psgo -format "rgroup"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "RGROUP" ]]
}

@test "NI" {
	run ./bin/psgo -format "%n"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "NI" ]]

	run ./bin/psgo -format "nice"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "NI" ]]
}

@test "PID header" {
	run ./bin/psgo -format "%p"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "PID" ]]

	run ./bin/psgo -format "pid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "PID" ]]
}

@test "ELAPSED header" {
	run ./bin/psgo -format "%t"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "ELAPSED" ]]

	run ./bin/psgo -format "etime"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "ELAPSED" ]]
}

@test "RUSER header" {
	run ./bin/psgo -format "%u"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "RUSER" ]]

	run ./bin/psgo -format "ruser"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "RUSER" ]]
}

@test "TIME header" {
	run ./bin/psgo -format "%x"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "TIME" ]]

	run ./bin/psgo -format "time"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "TIME" ]]
}

@test "TTY header" {
	run ./bin/psgo -format "%y"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "TTY" ]]

	run ./bin/psgo -format "tty"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "TTY" ]]
}

@test "VSZ header" {
	run ./bin/psgo -format "%z"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "VSZ" ]]

	run ./bin/psgo -format "vsz"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "VSZ" ]]
}

@test "CAPAMB header" {
	run ./bin/psgo -format "capamb"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "AMBIENT CAPS" ]]
}

@test "CAPINH header" {
	run ./bin/psgo -format "capinh"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "INHERITED CAPS" ]]
}

@test "CAPPRM header" {
	run ./bin/psgo -format "capprm"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "PERMITTED CAPS" ]]
}

@test "CAPEFF header" {
	run ./bin/psgo -format "capeff"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "EFFECTIVE CAPS" ]]
}

@test "CAPBND header" {
	run ./bin/psgo -format "capbnd"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "BOUNDING CAPS" ]]
}

@test "SECCOMP header" {
	run ./bin/psgo -format "seccomp"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "SECCOMP" ]]
}

@test "HPID header" {
	run ./bin/psgo -format "hpid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "HPID" ]]
	# host PIDs are only extracted with `-pid`
	[[ ${lines[1]} =~ "?" ]]
}

@test "HUSER header" {
	run ./bin/psgo -format "huser"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "HUSER" ]]
	# host users are only extracted with `-pid`
	[[ ${lines[1]} =~ "?" ]]
}

@test "HUID header" {
	run ./bin/psgo -format "huid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "HUID" ]]
	# host UIDs are only extracted with `-pid`
	[[ ${lines[1]} =~ "?" ]]
}

@test "HGROUP header" {
	run ./bin/psgo -format "hgroup"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "HGROUP" ]]
	# host groups are only extracted with `-pid`
	[[ ${lines[1]} =~ "?" ]]
}

@test "HGROUPS header" {
	run ./bin/psgo -format "hgroups"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "HGROUPS" ]]
	# host groups are only extracted with `-pid`
	[[ ${lines[1]} =~ "?" ]]
}


function is_labeling_enabled() {
	if [ -e /usr/sbin/selinuxenabled ] && /usr/sbin/selinuxenabled; then
			echo 1
			return
	fi
	echo 0
}

@test "LABEL header" {
	enabled=$(is_labeling_enabled)
	if [[ "$enabled" -eq 0 ]]; then
		skip "skip this test since labeling is not enabled."
	fi
	run ./bin/psgo -format "label"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "LABEL" ]]
}

@test "RSS header" {
	run ./bin/psgo -format "rss"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "RSS" ]]
}

@test "STATE header" {
	run ./bin/psgo -format "state"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "STATE" ]]
}

@test "STIME header" {
	run ./bin/psgo -format "stime"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ "STIME" ]]
}

@test "ALL header" {
	run ./bin/psgo -format "pcpu, group, groups, ppid, user, uid, args, comm, rgroup, nice, pid, pgid, etime, ruser, time, tty, vsz, capamb, capinh, capprm, capeff, capbnd, seccomp, hpid, huser, huid, hgroup, hgroups, rss, state"
	[ "$status" -eq 0 ]

	[[ ${lines[0]} =~ "%CPU" ]]
	[[ ${lines[0]} =~ "GROUP" ]]
	[[ ${lines[0]} =~ "GROUPS" ]]
	[[ ${lines[0]} =~ "PPID" ]]
	[[ ${lines[0]} =~ "USER" ]]
	[[ ${lines[0]} =~ "UID" ]]
	[[ ${lines[0]} =~ "COMMAND" ]]
	[[ ${lines[0]} =~ "COMMAND" ]]
	[[ ${lines[0]} =~ "RGROUP" ]]
	[[ ${lines[0]} =~ "NI" ]]
	[[ ${lines[0]} =~ "PID" ]]
	[[ ${lines[0]} =~ "ELAPSED" ]]
	[[ ${lines[0]} =~ "RUSER" ]]
	[[ ${lines[0]} =~ "TIME" ]]
	[[ ${lines[0]} =~ "TTY" ]]
	[[ ${lines[0]} =~ "VSZ" ]]
	[[ ${lines[0]} =~ "AMBIENT CAPS" ]]
	[[ ${lines[0]} =~ "INHERITED CAPS" ]]
	[[ ${lines[0]} =~ "PERMITTED CAPS" ]]
	[[ ${lines[0]} =~ "EFFECTIVE CAPS" ]]
	[[ ${lines[0]} =~ "BOUNDING CAPS" ]]
	[[ ${lines[0]} =~ "SECCOMP" ]]
	[[ ${lines[0]} =~ "HPID" ]]
	[[ ${lines[0]} =~ "HUSER" ]]
	[[ ${lines[0]} =~ "HUID" ]]
	[[ ${lines[0]} =~ "HGROUP" ]]
	[[ ${lines[0]} =~ "HGROUPS" ]]
	[[ ${lines[0]} =~ "RSS" ]]
	[[ ${lines[0]} =~ "STATE" ]]
}
