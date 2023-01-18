# syntax docker/dockerfile:1

# It's a comment. The very first comment must be after all directives
# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o asciiweb

FROM alpine:latest
LABEL progect-name="ascii-art-web-dockerize"
LABEL version="1.0.0"
LABEL description="creating a multi-stage builded docker for running ascii-art-web"
WORKDIR /app
COPY --from=builder /app/asciiweb .
COPY --from=builder /app/banners ./banners
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./asciiweb"]

