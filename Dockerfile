FROM devkitpro/devkitppc

VOLUME /rehover

ENV GOPATH=/go

# Install thirdparty ppc libs and other requirements
RUN apt-get update && \
	apt-get install -y --no-install-recommends git patch golang wget unzip cmake make && \
	git clone https://github.com/hoverguys/ppc-portlibs.git && \
	./ppc-portlibs/build.sh entityx && \
	rm -rf ./ppc-portlibs && \
	apt-get remove -y unzip wget patch && \
	apt-get autoremove -y && \
	apt-get clean

# Entrypoint
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]