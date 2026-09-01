# syntax=docker/dockerfile:1@sha256:ecfaec9ed6d810b56388c508f4121597bfbba70d41a6dfeee4d8cad5f295fc32
FROM golang:1.26.7-bookworm@sha256:e8c859f5632dcfde7b32d2012b4351728f6437930887c2f6a91ea242459e5514 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY cmd/thinkpixelmem/ cmd/thinkpixelmem/
COPY internal/ internal/
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -buildvcs=false \
    -ldflags='-s -w' -o /out/thinkpixelmem ./cmd/thinkpixelmem

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /out/thinkpixelmem /thinkpixelmem

USER 65532:65532
EXPOSE 8080
ENV TPMEM_HTTP_ADDRESS=0.0.0.0:8080
ENTRYPOINT ["/thinkpixelmem"]
