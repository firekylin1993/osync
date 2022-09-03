FROM harbor.myyun.org/inf/golang:1.17 as goBuilder
WORKDIR /go_build
COPY . /go_build
ARG VERSION_TAG
ARG COMMIT_ID
ENV VERSION_TAG ${VERSION_TAG:-0.0.0}
ENV GOPROXY=https://registry-mirrors.myyun.org/repository/goproxy/,https://goproxy.io,direct
ENV GONOPROXY=*.myyun.org
ENV GONOSUMDB=*.myyun.org
ENV GOPRIVATE=*.myyun.org
RUN go build -mod=vendor -ldflags "-X main.Version=$VERSION_TAG" -o /go/bin/ ./...
# debian release as the same as golang image
# set TimeZone as Asia/Shanghai
# set Local as zh-hans
FROM harbor.myyun.org/inf/debian:bullseye

ARG VERSION_TAG
ARG COMMIT_ID
ENV VERSION_TAG ${VERSION_TAG:-0.0.0}
EXPOSE 8080
COPY --from=goBuilder /go/bin/* /usr/local/bin/

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

ENTRYPOINT ["osync"]
CMD ["-conf", "/data/conf"]
