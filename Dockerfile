FROM scratch
LABEL maintainer="erikjo82+github@gmail.com"

ADD server /
CMD ["/server"]
