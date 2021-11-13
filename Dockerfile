FROM ubuntu:latest
COPY app_upgrade /root/server
EXPOSE 8080
EXPOSE 11451
ENV IS_DOCKER 1
CMD /root/server
