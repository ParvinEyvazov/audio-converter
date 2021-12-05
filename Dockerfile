FROM golang:1.16-alpine
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app .

FROM vimagick/youtube-dl
COPY --from=0 /app /app
ENTRYPOINT ["/app"]