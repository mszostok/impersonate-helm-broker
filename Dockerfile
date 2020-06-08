FROM alpine:latest

RUN apk --update add ca-certificates

COPY ./broker /bin/broker
COPY ./asset /bin/asset/

CMD ["/bin/broker"]
