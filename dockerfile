############################
# STEP 1 build executable binary
############################
FROM golang:1.22 as go-builder

# Create appuser
ENV USER=appuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download -x
RUN mkdir tmp && mkdir tmp/data 

COPY . .
RUN make docker-go-build

############################
# STEP 2 build FE files 
############################
FROM node:20 as node-builder

WORKDIR /node/src

COPY package.json package-lock.json ./
RUN npm install

COPY . .
RUN make docker-node-build

############################
# STEP 3 build a small image
############################
FROM gcr.io/distroless/static-debian12:debug

WORKDIR /app

ENV GIN_MODE=release

# Use an unprivileged user.
COPY --from=go-builder /etc/passwd /etc/passwd
COPY --from=go-builder /etc/group /etc/group
USER appuser:appuser

COPY --chown=appuser:appuser --from=go-builder /go/src/app/comment ./
COPY --chown=appuser:appuser --from=go-builder /go/src/app/tmp/data ./
COPY --chown=appuser:appuser --from=node-builder /node/src/app/assets/dist ./app/assets/dist
ENTRYPOINT ["./comment", "serve", "--https=false"] 
