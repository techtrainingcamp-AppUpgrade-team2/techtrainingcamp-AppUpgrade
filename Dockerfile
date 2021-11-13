FROM ubuntu:latest
COPY main /root/server
EXPOSE 9090
ENV IS_DOCKER 1
CMD /root/server
