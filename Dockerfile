FROM golang:1.19-alpine as builder
COPY --from=techdecaf/tasks /usr/local/bin/tasks /usr/local/bin/tasks
ADD . .
RUN tasks run build