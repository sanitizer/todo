# build app
FROM golang:alpine AS build-env
RUN apk --no-cache add git
ADD . /todo
RUN cd /todo && go build -o app

# run app
FROM alpine
WORKDIR /todo
COPY --from=build-env /todo/app /todo/
ENV DB_HOST=$DB_HOST
ENTRYPOINT ./app
