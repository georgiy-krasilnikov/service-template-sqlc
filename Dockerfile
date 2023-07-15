FROM golang:1.18-alpine as builder

ARG SERVICE_NAME
ENV BUILD_PATH="/tmp/${SERVICE_NAME}"

WORKDIR /
ADD . .

RUN go build -o ${BUILD_PATH} ./*.go


FROM alpine as runner

ARG SERVICE_NAME

ENV BIN_PATH="/usr/local/bin/${SERVICE_NAME}"
ENV SERVICE_PORT=${SERVICE_PORT}

COPY --from=builder "/tmp/${SERVICE_NAME}" ${BIN_PATH}
COPY --from=builder /migrations/ ./migrations

RUN chmod +x ${BIN_PATH}

CMD ${BIN_PATH}