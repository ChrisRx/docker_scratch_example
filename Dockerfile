FROM scratch
COPY tiny-server /
ENTRYPOINT ["/tiny-server"]
