# Couchbase 
#
# VERSION		1.0

FROM centos
MAINTAINER Nicolas Colomer ncolomer@ymail.com

ENV CB_VERSION		2.2.0
ENV CB_RELEASE_URL	http://packages.couchbase.com/releases
ENV CB_PACKAGE		couchbase-server-community_${CB_VERSION}_x86_64.rpm
ENV CB_USERNAME		Administrator
ENV CB_PASSWORD		couchbase

# Install couchbase
RUN yum install -y wget pkgconfig
RUN rpm --install $CB_RELEASE_URL/$CB_VERSION/$CB_PACKAGE
RUN ln -s /opt/couchbase/bin/couchbase-cli /usr/local/bin/

# Put start script
ADD sources/couchbase-start /usr/local/bin/

# See http://docs.couchbase.com/couchbase-manual-2.5/cb-install/#xdcr
EXPOSE 8091 8092 11210 11211

USER root
CMD couchbase-start 
