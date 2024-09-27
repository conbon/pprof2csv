PHONY: build

clean:
	rm -rf ./sample/*.prof
	rm -rf ./sample/*.txt
	rm -rf ./sample/*.csv

build:
	go build -o ./bin/pprof2csv pprof2csv.go
	go build -o ./bin/sample sample/bottleneck.go

run-sample:
	./bin/sample &

profile: run-sample
	sleep 1
	curl -s http://localhost:6111/debug/pprof/profile\?seconds\=10 -o sample/cpu.prof &
	curl -s http://localhost:6111/debug/pprof/heap\?seconds\=10 -o sample/heap.prof &

sample2csv:
	go tool pprof -text ./sample/heap.prof > sample/heap.txt
	go tool pprof -text ./sample/cpu.prof > sample/cpu.txt
	./bin/pprof2csv -input ./sample/heap.txt -output sample/heap.csv
	./bin/pprof2csv -input ./sample/cpu.txt -output sample/cpu.csv