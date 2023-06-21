FROM alpine:latest

COPY ./polygon /usr/bin/polygon
CMD ["polygon"]