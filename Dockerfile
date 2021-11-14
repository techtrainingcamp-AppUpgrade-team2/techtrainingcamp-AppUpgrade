FROM ubuntu:latest
COPY main /root/server
EXPOSE 8080
CMD /root/server
