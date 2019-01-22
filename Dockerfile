# Stage 1: Compile executable using the official Golang Alpine image.
FROM golang:alpine

WORKDIR /app

RUN apk --update add git openssh gcc musl-dev && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY . .
RUN go get -d -v && go build
RUN ls

# Stage 2: Copy executable to a lightweight image, and set the ENTRYPOINT
FROM alpine
COPY --from=0 /app/app /dnsrepeat
ENTRYPOINT [ "/dnsrepeat" ]
