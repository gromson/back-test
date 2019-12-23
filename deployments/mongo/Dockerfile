FROM mongo:latest
WORKDIR /
RUN apt-get update -y && \
    apt-get install apache2-utils -y
COPY ./wait-for-it.sh .
COPY ./deployments/mongo/create_admin.sh .
RUN chmod +x ./wait-for-it.sh
RUN chmod +x ./create_admin.sh
ENTRYPOINT ["./wait-for-it.sh", "mongo:27017", "--", "./create_admin.sh"]