FROM ubuntu:latest
COPY techtrainingcamp-AppUpgrade /root/server
EXPOSE 8080
EXPOSE 11451
ENV IS_DOCKER 1
CMD /root/server
