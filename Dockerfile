FROM alpine:3.4
ADD scores-srv /scores-srv
ENTRYPOINT [ "/scores-srv" ]
