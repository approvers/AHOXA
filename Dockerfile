FROM golang:1.14.4-alpine3.12 as build
RUN mkdir /src
COPY . /src
WORKDIR /src
RUN go build -a -tags netgo -installsuffix netgo -o AHOXA


FROM alpine:3.12
COPY --from=build /src/AHOXA /bot/
WORKDIR /bot/
ENTRYPOINT [ "/bot/AHOXA" ]
