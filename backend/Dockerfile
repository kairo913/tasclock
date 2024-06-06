FROM golang:1.22.2-bookworm

RUN go install github.com/cosmtrek/air@latest

CMD [ "go", "mod", "tidy" ]