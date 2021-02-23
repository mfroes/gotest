FROM golang:1.16.0 as build
ARG COMMIT_SHA=""
ARG VERSION=""
ARG BUILD_TIME=""
ARG APP_DESCRIPTION=""

WORKDIR /build/src
COPY . /build/src/
# Global metadata
LABEL app_name="GoTest"
LABEL maintaner="Marcelo Froes <mfroes@gmail.com>"
# Setting up docker metadata from the version file
LABEL commit_sha=${COMMIT_SHA} version=${VERSION} app_description=${APP_DESCRIPTION} build_time=${BUILD_TIME}

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/bin/server \
  -ldflags="-X github.com/mfroes/gotest/main/pkg.BuildVersion=${VERSION} \
  -X 'github.com/mfroes/gotest/main/pkg.BuildSHA=${COMMIT_SHA}' \
  -X 'github.com/mfroes/gotest/main/pkg.BuildDescription=${APP_DESCRIPTION}' \
  -X 'github.com/mfroes/gotest/main/pkg.BuildTime=${BUILD_TIME}'" main.go

# HERE STARTS THE DISTRIBUTION DOCKER IMAGE
FROM scratch as dist
# Arguments to set the metadata
ARG COMMIT_SHA=""
ARG VERSION=""
ARG BUILD_TIME=""
ARG APP_DESCRIPTION=""

# Global metadata
LABEL app_name="GoTest"
LABEL maintaner="Marcelo Froes <mfroes@gmail.com>"
# Setting up docker metadata from the version file
LABEL commit_sha=${COMMIT_SHA} version=${VERSION} app_description=${APP_DESCRIPTION} build_time=${BUILD_TIME}

COPY --from=build /build/bin/server /server
EXPOSE 8090
ENTRYPOINT ["/server"]
