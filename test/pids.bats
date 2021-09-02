#!/usr/bin/env bats -t

@test "Get process information of one process only (init)" {
	pid="1"
	run ./bin/psgo -pids "$pid" -format "user,pid,ppid"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ ^USER\ +PID\ +PPID$ ]]
	[[ ${lines[1]} =~ ^root\ +$pid\ +0$ ]]
}

@test "Get process information of a set of running containers" {
	nCtrs=5
	pidsList=()
	ctridList=()
	for (( i=1; i<=$nCtrs; i++ )); do
		ctridList[$i]="$(docker run -d busybox sleep 60)"
		pidsList[$i]="$(docker inspect --format '{{.State.Pid}}' ${ctridList[$i]})"
	done

	pids="$(tr ' ' ',' <<< ${pidsList[@]})"

	run ./bin/psgo -pids "$pids" -format "pid,comm"

	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ ^PID\ +COMMAND$ ]]
	for (( i=1; i<=$nCtrs; i++ )); do
		[[ ${lines[$i]} =~ ^${pidsList[$i]}\ +sleep$ ]]
		docker rm -f ${ctridList[$i]}
	done
}

@test "Return error on both --pid and --pids options" {
	run ./bin/psgo -pid 1 --pids 1,2,3
	! [ "$status" -eq  0]
}

@test "Process information with --pids vs all processes" {
	nCtrs=5
	pidsList=()
	ctridList=()
	for (( i=1; i<=$nCtrs; i++ )); do
		ctridList[$i]="$(docker run -d busybox sleep 60)"
		pidsList[$i]="$(docker inspect --format '{{.State.Pid}}' ${ctridList[$i]})"
	done

	pids="$(tr ' ' ',' <<< ${pidsList[@]})"

	run ./bin/psgo --pids "$pids" -format "user,group,pid,ppid,tty,nice,capeff"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ ^USER\ +GROUP\ +PID\ +PPID\ +TTY\ +NI\ +EFFECTIVE\ CAPS$ ]]
	pidsLines=${lines[@]}

	run ./bin/psgo -format "user,group,pid,ppid,tty,nice,capeff"
	[ "$status" -eq 0 ]
	[[ ${lines[0]} =~ ^USER\ +GROUP\ +PID\ +PPID\ +TTY\ +NI\ +EFFECTIVE\ CAPS$ ]]

	for (( i=1; i<=$nCtrs; i++ )); do
		docker rm -f ${ctridList[$i]}
	done
}
