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

	run sudo ./bin/psgo -pid $PID -format "pid, args"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   COMMAND" ]]
	[[ ${lines[1]} =~ "1     sleep 100" ]]

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
