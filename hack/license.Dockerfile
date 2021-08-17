# syntax=docker/dockerfile:1.3

# Copyright 2021 docgen authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG GO_VERSION
ARG ADDLICENSE_VERSION="d43bb61fdfdafb29f4b1add4b849c5bfe4eeb497"
ARG LICENSE_ARGS="-c docgen -l apache"
ARG LICENSE_FILES=".*\(Dockerfile\|\.go\|\.hcl\|\.sh\)"

FROM golang:${GO_VERSION}-alpine AS base
WORKDIR /src
RUN apk add --no-cache cpio findutils git
ENV CGO_ENABLED=0
ARG ADDLICENSE_VERSION
RUN go install github.com/google/addlicense@${ADDLICENSE_VERSION}

FROM base AS set
ARG LICENSE_ARGS
ARG LICENSE_FILES
RUN --mount=type=bind,target=.,rw \
  find . -regex "${LICENSE_FILES}" | xargs addlicense ${LICENSE_ARGS} \
  && mkdir /out \
  && find . -regex "${LICENSE_FILES}" | cpio -pdm /out

FROM scratch AS update
COPY --from=set /out /

FROM base AS validate
ARG LICENSE_ARGS
ARG LICENSE_FILES
RUN --mount=type=bind,target=. \
  find . -regex "${LICENSE_FILES}" | xargs addlicense -check ${LICENSE_ARGS}