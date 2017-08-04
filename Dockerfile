FROM alpine:3.4
ADD qrcode-srv /qrcode-srv
ENTRYPOINT [ "/qrcode-srv" ]
