
FROM ubuntu:22.04


ENV DEBIAN_FRONTEND=noninteractive


RUN apt-get update && \
    apt-get install -y \
        ca-certificates \
        wget \
	      net-tools\
        pdns-backend-sqlite3 \
        pdns-backend-bind && \
        update-ca-certificates && \
        apt-get clean && \
        rm -rf /var/lib/apt/lists/*


EXPOSE 53/udp
EXPOSE 53/tcp

EXPOSE 8081


CMD ["pdns_server", "--daemon=no", "--guardian=no"]
