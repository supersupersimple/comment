############################
# STEP 1 build executable binary
############################
FROM golang:1.22 as go-builder

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
FROM gcr.io/distroless/static-debian12:nonroot

ENV GIN_MODE=release

COPY --from=go-builder --chown=nonroot:nonroot /go/src/app/comment ./
COPY --from=go-builder --chown=nonroot:nonroot /go/src/app/tmp/data ./
COPY --from=node-builder --chown=nonroot:nonroot /node/src/app/assets/dist ./app/assets/dist
VOLUME ["/home/nonroot/data"]
CMD ["./comment", "serve", "--https=false"] 
