FROM alpine:3.11.6
COPY main /app/
EXPOSE 8000
WORKDIR /app
CMD ["/app/main"]

