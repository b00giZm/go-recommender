FROM scratch
MAINTAINER team-benutzer@chefkoch.de

COPY build/recommender /recommender

ENTRYPOINT ["/recommender"]
