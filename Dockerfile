FROM alpine:3.5
MAINTAINER Hector Rios

#ENV SOURCES /go/src/json_marshalling/
#COPY . ${SOURCES}
#RUN cd ${SOURCES} && CGO_ENABLED=0 go install

COPY ./json_marshalling /app/json_marshalling
RUN chmod +x /app/json_marshalling

ENV PORT 8080
EXPOSE 8080

#ENTRYPOINT json_marshalling
ENTRYPOINT /app/json_marshalling


