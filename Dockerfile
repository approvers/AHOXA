FROM golang:1.14.4-alpine3.12 as build
RUN mkdir /src
COPY . /src
WORKDIR /src
RUN go build -o change-status-go


FROM alpine:3.12
COPY --from=build /src/change-status-go /bot/
WORKDIR /bot/
ENTRYPOINT [ "/bot/change-status-go" ]
