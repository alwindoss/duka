FROM golang:1.26.2 AS builder
# Install dependencies
RUN apt-get update && apt-get install -y \
    curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa \
    fonts-droid-fallback lib32stdc++6 python3 make
RUN curl https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_3.41.6-stable.tar.xz
RUN tar -xf flutter_linux_3.41.6-stable.tar.xz -C /usr/local/
# Set Flutter path
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"
# Enable web and build
RUN flutter doctor -v
RUN flutter config --enable-web

WORKDIR /go/src/app
COPY . .
RUN make package

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian13
WORKDIR /root/
COPY --from=builder /go/src/app/bin/duka .
CMD [ "./duka" ]