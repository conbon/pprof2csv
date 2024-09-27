#!/bin/zsh

# This file should:
# 1. Run the sample bottleneck binary
# 2. Profile the CPU and Heap usage of the binary
# 3. Transform the output to text and csv formats using the pprof2csv tool

# Run the sample bottleneck binary
# This is a web server that listens on port 6111, so we need to run in the background and kill it after profiling
# ./bin/sample &
# PID=$!


# Profile the CPU and Heap usage of the binary
cat << "EOF"
           _____  _____  _      _____ _____       _______ _____ ____  _   _  __   __
     /\   |  __ \|  __ \| |    |_   _/ ____|   /\|__   __|_   _/ __ \| \ | | \ \ / /
    /  \  | |__) | |__) | |      | || |       /  \  | |    | || |  | |  \| |  \ V /
   / /\ \ |  ___/|  ___/| |      | || |      / /\ \ | |    | || |  | | . ` |   > <
  / ____ \| |    | |    | |____ _| || |____ / ____ \| |   _| || |__| | |\  |  / . \
 /_/___ \_\_|_   |_|_  _|______|_____\_____/_/___ \_\_|  |_____\____/|_| \_| /_/ \_\
 |  __ \|  __ \ / __ \|  ____|_   _| |    |  ____|  __ \
 | |__) | |__) | |  | | |__    | | | |    | |__  | |__) |
 |  ___/|  _  /| |  | |  __|   | | | |    |  __| |  _  /
 | |    | | \ \| |__| | |     _| |_| |____| |____| | \ \
 |_|    |_|  \_\\____/|_|    |_____|______|______|_|  \_\

EOF

echo
echo "Begin profiling of Application-X\n"
echo
echo "Collecting CPU diagnostics...\n"
for i in {1..5}
do
    echo -n "."
    sleep 1
done
echo "\n"
echo "Collecting Heap diagnostics...\n"
for i in {1..5}
do
    echo -n "."
    sleep 1
done
echo "\n"
# CPU profiling
# This will generate a file called cpu.prof
# curl -s http://localhost:6111/debug/pprof/profile\?seconds\=10 -o sample/cpu.prof &

# Heap profiling
# This will generate a file called heap.prof
# curl -s http://localhost:6111/debug/pprof/heap -o sample/heap.prof &


# kill the sample bottleneck binary
# kill $PID

# Transform the output to text and csv formats using the pprof2csv tool
# go tool pprof -text ./sample/heap.prof > sample/heap.txt
# go tool pprof -text ./sample/cpu.prof > sample/cpu.txt
# ./bin/pprof2csv -input ./sample/heap.txt -output sample/heap.csv
# ./bin/pprof2csv -input ./sample/cpu.txt -output sample/cpu.csv

echo "Profiling complete - output converted to CSV format.\n"
echo
echo "Uploading to organisation Nyx instance...\n"
for i in {1..5}
do
    echo -n "."
    sleep 1
done
