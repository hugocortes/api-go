FROM alpine
COPY paprika-api .
ENTRYPOINT [ "/paprika-api" ]
