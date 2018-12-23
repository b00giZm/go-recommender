FROM debian:jessie
MAINTAINER b00gizm@gmail.com

RUN apt-get update \
    && apt-get upgrade \
    && apt-get install -yq \
        beanstalkd \
    && apt-get autoremove \
    && apt-get clean

CMD ["beanstalkd", "-l", "0.0.0.0", "-p", "11300"]
