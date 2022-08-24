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

	run sudo ./bin/psgo -pids $PID -join
	[ "$status" -eq 0 ]
	[[ ${lines[1]} =~ "sleep" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and format" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, group, args"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   GROUP   COMMAND" ]]
	[[ ${lines[1]} =~ "1     root    sleep 100" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and check capabilities" {
	ID="$(docker run --privileged -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, capeff"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   EFFECTIVE CAPS" ]]
	# FIXME: we don't get "full" anymore running against Docker and Podman
	# seems to report some unknown ones.  I didn't see new capabilities in
	# the kernel source, so there may be more to it.
#	[[ ${lines[1]} =~ "1     full" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and check seccomp mode" {
	# Run a privileged container to force seecomp to "disabled" to avoid
	# hiccups in Travis.
	ID="$(docker run -d --privileged alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID --join -format "pid, seccomp"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   SECCOMP" ]]
	[[ ${lines[1]} =~ "1     disabled" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and extract host PID" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, hpid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   HPID" ]]
	[[ ${lines[1]} =~ "1     $PID" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and extract effective host user ID" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, huser"
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

	ID="$(sudo podman run -d --uidmap=0:300000:70000 --gidmap=0:100000:70000 alpine sleep 100)"
	PID="$(sudo podman inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, user, huser, group, hgroup"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   USER   HUSER    GROUP   HGROUP" ]]
	[[ ${lines[1]} =~ "1     root   300000   root    100000" ]]

	sudo podman rm -f $ID
}

@test "Join namespace of a Docker container and extract effective host group ID" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, hgroup"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   HGROUP" ]]
	[[ ${lines[1]} =~ "1     root" ]]

	docker rm -f $ID
}

@test "Join namespace of a Docker container and check the process state" {
	ID="$(docker run -d alpine sleep 100)"
	PID="$(docker inspect --format '{{.State.Pid}}' $ID)"

	run sudo ./bin/psgo -pids $PID -join -format "pid, state"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   STATE" ]]
	[[ ${lines[1]} =~ "1     S" ]]

	docker rm -f $ID
}

@test "Run Podman pod and check for redundant entries" {
	enabled=$(is_podman_available)
	if [[ "$enabled" -eq 0 ]]; then
		skip "skip this test since Podman is not available."
	fi

	POD_ID="$(sudo podman pod create)"
	ID_1="$(sudo podman run --pod $POD_ID -d alpine sleep 111)"
	PID_1="$(sudo podman inspect --format '{{.State.Pid}}' $ID_1)"
	ID_2="$(sudo podman run --pod $POD_ID -d alpine sleep 222)"
	PID_2="$(sudo podman inspect --format '{{.State.Pid}}' $ID_2)"

	# The underlying idea is that is that we had redundant entries if
	# the detection of PID namespaces wouldn't work correctly.
	run sudo ./bin/psgo -pids "$PID_1,$PID_2" -join -format "pid, args"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} == "PID   COMMAND" ]]
	[[ ${lines[1]} =~ "1     sleep 111" ]]
	[[ ${lines[2]} =~ "1     sleep 222" ]]
	[[ ${lines[3]} = "" ]]

	sudo podman rm -f $ID_1 $ID_2
	sudo podman pod rm $POD_ID
}

@test "Test fill-mappings" {
	if [[ ! -z "$TRAVIS" ]]; then
		skip "Travis doesn't like this test"
	fi

	run unshare -muinpfr --mount-proc true
	if [[ "$status" -ne 0 ]]; then
		skip "unshare doesn't support all the needed options"
	fi

	unshare -muinpfr --mount-proc sleep 20 &

	PID=$(echo $!)
	run nsenter --preserve-credentials -U -t $PID ./bin/psgo -pids $PID -join -fill-mappings -format huser
	kill -9 $PID
	[ "$status" -eq 0 ]
	[[ ${lines[0]} != "root" ]]
}
