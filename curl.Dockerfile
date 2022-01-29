# experimental
FROM alpine:latest AS alpine
RUN set -x \
 && apk --no-cache add curl 

#FROM scratch
#COPY --from=alpine /usr/bin/curl /bin/curl
#ENTRYPOINT ["/usr/bin/curl"]
#CMD ["localhost"]