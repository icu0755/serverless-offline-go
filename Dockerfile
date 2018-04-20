FROM lambci/lambda:build-go1.x
RUN yum install -y -t ImageMagick-devel || true
