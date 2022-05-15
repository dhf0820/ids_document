FROM alpine:latest

ADD document_arm64 ./document
RUN mkdir ./config
RUN mkdir /root/tmp_images
RUN mkdir /root/data
RUN mkdir /root/doc_data
#ADD ./config/core.json ./config/
ADD .env ./.env.document
#EXPOSE 50051
EXPOSE 9200
ENTRYPOINT ["./document"]#