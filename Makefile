build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/sign-in functions/sign-in.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/authorizer functions/authorizer.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/list-pools functions/list-pools.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/list-users functions/list-users.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/user-enabled functions/user-enabled.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/user-details functions/user-details.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update-user-attributes functions/update-user-attributes.go
