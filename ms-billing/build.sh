#!/bin/bash

cd /home/ossan/Projects/go-debug-vsc/ms-billing
echo 0 | tee /proc/sys/kernel/yama/ptrace_scope
go build -modfile=../go.mod -gcflags=all="-N -l"
PORT=8081 ./ms-billing