# **********************
# docker build --no-cache --progress=plain -t golang_server:v1.0.1 -f docker/go/Dockerfile .
# docker build --progress=plain -t golang_server:v1.0.1 -f docker/go/Dockerfile .
# docker run golang_server:v1.0.1
# **********************
# 1. Build Golang Application
FROM golang:1.16.4 as build_go_app
COPY . /build
WORKDIR /build
RUN apt-get update && apt-get install -y python3-all-dev
RUN go build -ldflags '-w -s' -o server

# 2. Get python-model dist
FROM alpine:3.11 as get_python_model_dist
ENV PYTHON_DIST_FILE_NAME='python-model-1.0.1.tar.gz'
ENV PYTHON_DIST_FILE_IN_PYTHON_MODEL_IMAGE=/build/__dist__/${PYTHON_DIST_FILE_NAME}
ENV PYTHONPATH=/build/pythonsrc/
COPY --from=python-model:v1.0.1 ${PYTHON_DIST_FILE_IN_PYTHON_MODEL_IMAGE} ${PYTHONPATH}
WORKDIR ${PYTHONPATH}
RUN tar -xzvf ${PYTHON_DIST_FILE_NAME} --strip-components 1 -C .

# 3. Combine golang app + python-mode dist
FROM python:3.7 as deployable
WORKDIR /home/goserver/
COPY --from=build_go_app /build/server .
COPY --from=get_python_model_dist /build/pythonsrc/ ./py_src/
ENV PYTHONPATH=/home/goserver/py_src/
ENV PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:/usr/local/lib/pkgconfig/
ENTRYPOINT ["/home/goserver/server"]

