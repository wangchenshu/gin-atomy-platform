FROM golang:latest

WORKDIR /web/atomy
COPY . .
RUN go build
CMD ["./gin-atomy-platform"]
ENTRYPOINT ["./bootstrap.sh"]
