FROM golang:1.24-alpine AS BUILDER

WORKDIR /app

RUN apk add --no-cache git
RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY . ./
RUN task build

FROM scratch AS RUNNER

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

LABEL org.opencontainers.image.created=$BUILD_DATE
LABEL org.opencontainers.image.revision=$VCS_REF
LABEl org.opencontainers.image.version=$VERSION
LABEL org.opencontainers.image.title="Datastar-Playground"
LABEL org.opencontainers.image.description="Datastar Playground image serving an OSRS goal planner."
LABEL org.opencontainers.image.source="https://github.com/johnfarrell/datastar-playground"
LABEL org.opencontainers.image.authors="John Farrell (johnjfarrell20@gmail.com)"

COPY --from=BUILDER /app/build/dsp /dsp

EXPOSE 8080

CMD [ "/dsp", "start" ]