FROM golang:1.14.4-alpine3.12 as build
RUN mkdir /src
COPY . /src
WORKDIR /src
RUN go build -a -tags netgo -installsuffix netgo -o change-status-go


FROM scratch
COPY --from=build /src/change-status-go /bot/
WORKDIR /bot/
ENTRYPOINT [ "/bot/change-status-go" ]
