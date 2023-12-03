FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go docker cmd/main/main.go

FROM ubuntu:latest
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/config/user /app/config/user
COPY --from=build /app/config/auth /app/config/auth
COPY --from=build /app/config/payment /app/config/payment
CMD ["./main"]









