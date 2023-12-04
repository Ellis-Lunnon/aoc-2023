run:
	go run days/$(day)/$(day).go

start:
	mkdir days/$(day)
	echo "package main\n\n\nfunc Part1() {\n\n}\nfunc Part2() {\n\n}\nfunc main() {\n\n}" > days/$(day)/$(day).go
