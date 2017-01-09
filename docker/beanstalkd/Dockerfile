FROM debian:jessie
MAINTAINER team-benutzer@chefkoch.de

RUN apt-get update \
    && apt-get upgrade \
    && apt-get install -yq \
        beanstalkd \
    && apt-get autoremove \
    && apt-get clean

CMD ["beanstalkd", "-l", "0.0.0.0", "-p", "11300"]
