# ---- build stage ----
FROM golang:1.23 AS build
WORKDIR /app

# cache deps dulu
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy source
COPY . .

# penting: build SATU paket (.)
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v -trimpath -buildvcs=false -o /out/server .

# ---- runtime ----
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/server .
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./server"]
