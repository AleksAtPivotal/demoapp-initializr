FROM golang:latest as builder
WORKDIR /go/src/github.com/adnanh/webhook
RUN cd /go/src/github.com/adnanh ;\
    git clone https://github.com/adnanh/webhook.git ; \
    cd webhook ; \
    git checkout tags/2.6.8 ; \
    CGO_ENABLED=0 GOOS=linux go build . ;\
    mkdir /app && \ 
    wget https://github.com/concourse/concourse/releases/download/v4.2.1/fly_linux_amd64 -q && \
    chmod +x fly_linux_amd64 && \
    mv ./fly_linux_amd64 /app/fly && \
    mv ./webhook /app/webhook

FROM alpine:latest
RUN mkdir /app 
WORKDIR /app
COPY --from=builder /app .
COPY ./ . 
CMD ["/app/webhook", "-verbose", "-urlprefix=", "-port=8080"]
EXPOSE 8080