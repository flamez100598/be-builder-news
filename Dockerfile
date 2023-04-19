FROM ubuntu:22.04
ENV JWT_SERVICE_SECRET=VMQnMzKSpKg8FP7JKQEU
ENV MYSQL_ADDR=localhost:3306
ENV MYSQL_DATABASE=newsdb
ENV MYSQL_USER=nuser
ENV MYSQL_PASS=nujdh72#5t
WORKDIR /app
RUN mkdir config
COPY deploy/config/config.yml config/config.yml
COPY ./app_main /app/app_main
EXPOSE 8084
ENTRYPOINT ["/app/app_main"]
