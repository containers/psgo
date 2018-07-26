#!/usr/bin/env bats -t

@test "Join namespace of a Docker container" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID
	[ "$status" -eq 0 ]
	[[ ${lines[1]} =~ "sleep" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and format" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, group, args"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   GROUP   COMMAND" ]]
	[[ ${lines[1]} =~ "1     root    sleep 100" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and check capabilities" {
	ID="$(docker run --privileged -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, capeff"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   CAPABILITIES" ]]
	[[ ${lines[1]} =~ "1     full" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and check seccomp mode" {
	# Run a privileged container to force seecomp to "disabled" to avoid
	# hiccups in Travis.
	ID="$(docker run -d --privileged alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, seccomp"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   SECCOMP" ]]
	[[ ${lines[1]} =~ "1     disabled" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and extract host PID" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, hpid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   HPID" ]]
	[[ ${lines[1]} =~ "1     $PID" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and extract effective host user ID" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, huser"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   HUSER" ]]
	[[ ${lines[1]} =~ "1     root" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and extract effective host group ID" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, hgroup"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   HGROUP" ]]
	[[ ${lines[1]} =~ "1     root" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and check the process state" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, state"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   STATE" ]]
	[[ ${lines[1]} =~ "1     S" ]]

	docker rm -f $ID
}
