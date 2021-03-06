# build stage
FROM sevendollar/golang AS build-env
COPY main.go /go/src/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app .

# final stage
FROM busybox
COPY --from=build-env /go/src/app /
EXPOSE 80
ENTRYPOINT /app