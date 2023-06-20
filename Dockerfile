FROM alpine:latest

COPY ./polygon /usr/bin
CMD ["polygon"]