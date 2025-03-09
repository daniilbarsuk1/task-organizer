FROM ubuntu:24.04
WORKDIR /home
COPY src ./
EXPOSE 8080
WORKDIR /home/main
ENTRYPOINT [ "./main" ]