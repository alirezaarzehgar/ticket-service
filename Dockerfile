FROM golang:1.21.3
WORKDIR /app
COPY . .
RUN go build -o server
ENTRYPOINT [ "bash", "-c" ]
CMD [ "./server" ]
