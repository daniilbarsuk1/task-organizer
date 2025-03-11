FROM alpine:3.21.3
WORKDIR /home
RUN apk add --no-cache libc6-compat 
COPY src ./
EXPOSE 8080
WORKDIR /home/main
ENTRYPOINT [ "./main" ]