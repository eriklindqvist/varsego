FROM scratch
LABEL maintainer="erikjo82+github _at_ gmail.com"

ADD server /
CMD ["/server"]
