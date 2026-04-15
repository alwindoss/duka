.PHONY: dev build run clean setup package
dev:
	go tool task dev --parallel
setup:
	go tool task setup

build:
	go tool task build

run:
	go tool task run

clean:
	go tool task clean

package:
	go tool task package