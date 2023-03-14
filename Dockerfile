FROM golang:1.18
ENV PORT 3000
ENV DB_PASSWD PoxlRbrdppQoqw1q
ENV CUR_DB testing
WORKDIR Documentos/sports/sports
COPY . .
RUN go build -o bin/server main.go
CMD ["./bin/server"]