#!/usr/bin/env bats -t

function is_podman_available() {
	if podman help >> /dev/null; then
			echo 1
			return
	fi
	echo 0
}

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
	[[ ${lines[0]} == "PID   EFFECTIVE CAPS" ]]
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

@test "Join namespace of a Podman container and extract pid, {host,}user and group with {g,u}idmap" {
	enabled=$(is_podman_available)
	if [[ "$enabled" -eq 0 ]]; then
		skip "skip this test since Podman is not available."
	fi

	ID="$(podman run -d --uidmap=0:300000:70000 --gidmap=0:100000:70000 alpine sleep 100)"
	PID="$(podman inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pid $PID -format "pid, user, huser, group, hgroup"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   USER   HUSER    GROUP   HGROUP" ]]
	[[ ${lines[1]} =~ "1     root   300000   root    100000" ]]

	podman rm -f $ID
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
