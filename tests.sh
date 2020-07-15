#!/bin/bash
go run main.go &
export COUNT_URL='http://localhost:443/count'

cd service || exit
echo "************************ testing service ************************************"
go test -v -coverprofile=cover.out && go tool cover -func=cover.out
cd ..

# There's no coverage report bc I'm handling this more like an intergration test.
# With a more complex API I'd do both. We're really just testing validation here.
cd api || exit
echo "************************ testing api ************************************"
go test -v
cd ..

echo "************************ testing main ************************************"
go test -v -coverprofile=cover.out && go tool cover -func=cover.out

