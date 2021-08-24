FROM ubuntu:latest
USER root
RUN mkdir /root/gs2migrations3 && \
    mkdir -p /root/.config/rclone && \
    apt-get update && \
    apt-get install -y curl && \
    apt-get install -y unzip && \
    curl https://rclone.org/install.sh | bash
COPY ./build/bin/gs2migrations3/gs2migrations3 /root/gs2migrations3/gs2migrations3
EXPOSE 8083
ENTRYPOINT ["/root/gs2migrations3/gs2migrations3"]