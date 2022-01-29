FROM busybox
ENTRYPOINT ["/bin/ping", "-c", "3"]
CMD ["localhost"]