run:
	go run days/$(day)/$(day).go

start:
	go run maketemplate.go -d $(day)
