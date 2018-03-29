FROM alpine:3.2
ADD db-srv-srv /db-srv-srv
ENTRYPOINT [ "/db-srv-srv" ]
