FROM alpine:3.18.5

RUN apk add --no-cache bash
RUN adduser --system --disabled-password cxuser
USER cxuser

COPY cx /app/bin/cx

ENTRYPOINT ["/app/bin/cx"]
