FROM alpine:latest AS alpine
WORKDIR /bin
RUN set -x \
 && apk --no-cache add curl 

FROM scratch
#COPY --from=alpine /bin/command /bin/command
#ENTRYPOINT ["/bin/command"]
#CMD ["option"]
