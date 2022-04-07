module github.com/interviews-go/event-hub

go 1.17

replace (
	github.com/interviews-go/event-hub latest => ./ latest
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/pkg/errors v0.9.1
)

