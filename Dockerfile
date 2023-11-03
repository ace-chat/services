FROM centos:latest
LABEL author=alex

# Build Args
ARG LOG_DIR=/sites/logs

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log GO111MODULE=on

# Create Log Directory
RUN mkdir -p ${LOG_DIR} \
    && cd /etc/yum.repos.d/ \
    && sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* \
    && sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-* \
    && yum update -y \
    && yum install -y wget tar \
    && cd /usr/local \
    && wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz \
    && tar zxvf go1.21.3.linux-amd64.tar.gz \
    && rm -rf go1.21.3.linux-amd64.tar.gz

WORKDIR /sites

COPY . .

RUN /usr/local/go/bin/go mod tidy \
    && /usr/local/go/bin/go build -o ace \
    && mv ./ace /usr/local/bin \
    && rm -rf ./* && rm -rf /usr/local/go && yum clean all

EXPOSE 9000

CMD ["/usr/local/bin/ace"]