FROM debian

ENV DEVKITPRO=/opt/devkitpro \
	DEVKITPPC=/opt/devkitpro/devkitPPC

# Install devKitPro
RUN apt-get update && \
	apt-get install -y --no-install-recommends apt-utils ca-certificates && \
	apt-get install -y --no-install-recommends gpg libxml2 wget xz-utils && \
	wget https://github.com/devkitPro/pacman/releases/download/devkitpro-pacman-1.0.1/devkitpro-pacman.deb -P /home/root/ && \
	dpkg --install /home/root/devkitpro-pacman.deb && \
	dkp-pacman --noconfirm -S devkit-env devkitPPC gamecube-tools general-tools libfat-ogc libogc && \
	dkp-pacman --noconfirm -Scc && \
	rm /home/root/devkitpro-pacman.deb && \
	apt-get remove -y gpg libxml2 wget xz-utils && \
	apt-get autoremove -y && \
	apt-get clean