FROM alpine:latest

WORKDIR /root/

COPY .bin/ .

EXPOSE 3000

ENTRYPOINT [ "./main" ]
