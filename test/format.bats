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

@test "ALL header" {
	run ./bin/psgo -format "pcpu, group, ppid, user, args, comm, rgroup, nice, pid, pgid, etime, ruser, time, tty, vsz"
	[ "$status" -eq 0 ]

	[[ ${lines[0]} =~ "%CPU" ]]
	[[ ${lines[0]} =~ "GROUP" ]]
	[[ ${lines[0]} =~ "PPID" ]]
	[[ ${lines[0]} =~ "USER" ]]
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
}
