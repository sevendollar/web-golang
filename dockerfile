# build stage
FROM sevendollar/golang AS build-env
ENV TZ=Asia/Taipei
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime
RUN echo $TZ > /etc/timezone
COPY main.go /go/src/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app .

# final stage
FROM alpine
COPY --from=build-env /go/src/app /go/src/
EXPOSE 80
ENTRYPOINT /go/src/app