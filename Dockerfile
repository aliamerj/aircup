FROM node:20 as client-build
WORKDIR /build

COPY ./client/package.json ./client/yarn.lock .
RUN yarn install --frozen-lockfile

COPY ./client .
RUN yarn build

FROM golang:1.22.2-alpine3.19 as build

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy the modules files
COPY go.mod .
COPY go.sum .

# Download the modules
RUN go mod download

# Copy rest fo the code
COPY . .

# Copt the frontend build into the expected folder
COPY --from=client-build /build/dist ./client/dist

RUN CGO_ENABLED=0 ENV=prod go build -buildvcs=false -o ./bin/aircup ./main.go

FROM alpine:3.19

COPY --from=build /build/bin/aircup /usr/bin/aircup

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run the executable
CMD /usr/bin/aircup
