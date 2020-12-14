FROM scratch

COPY cmd/cmd .

EXPOSE 8099

ENTRYPOINT [ "./cmd" ]