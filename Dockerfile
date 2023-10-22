FROM golang:1.19 as builder

# Set destination for COPY
WORKDIR /app

RUN #apt-get update && apt-get install libc6 -y

# Download and build swisseph
RUN wget https://github.com/aloistr/swisseph/archive/refs/tags/v2.10.03.tar.gz
RUN tar -zxvf v2.10.03.tar.gz
RUN cd swisseph-2.10.03 && \
    echo 'build-libswe.so: $(SWEOBJ)' >> Makefile && \
    echo '\t$(CC) -shared -o libswe.so $(SWEOBJ) -lm -ldl' >> Makefile

RUN cd swisseph-2.10.03 && \
    make build-libswe.so && \
    cp libswe.so /usr/local/lib/

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Build
RUN go build -o app -ldflags '-libgcc' .

FROM golang:1.19 as release

EXPOSE 3000
# Copy the Go application binary from the builder image
ENV LD_LIBRARY_PATH=/usr/local/lib
COPY --from=builder /app/app /app
COPY --from=builder /app/swisseph-2.10.03/libswe.so /usr/local/lib/libswe.so

# Run
CMD ["/app"]