FROM golang:1.17.10-alpine3.15 AS build_stage

# setup environment
ARG work_dir=/app
ARG project=event-scheduler
RUN mkdir -p ${work_dir}/src
WORKDIR ${work_dir}/src

# build code
ENV GO111MODULE=on
ENV GOPATH=${work_dir}

# In build process, makefile will copy your ${HOME}/.ssh into project folder as ssh_keys.
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download -x
COPY . ${work_dir}/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/${project}/main cmd/${project}/main.go

# In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine:3.16.0
ARG work_dir=/app
ARG project=event-scheduler
WORKDIR ${work_dir}
RUN mkdir -p ${work_dir} && mkdir -p ${work_dir}/configs

# copy binary into image
COPY --from=build_stage ${work_dir}/src/cmd/${project}/main ${work_dir}/main

# entrypoint
COPY build/${project}/entrypoint.sh ${work_dir}/entrypoint.sh
COPY configs/config.yaml ${work_dir}/config.yaml

ENV WORK_DIR ${work_dir}
ENTRYPOINT ["sh", "entrypoint.sh"]
