FROM quay.io/brianredbeard/corebox

ADD shopapi /bin/shopapi
ADD entrypoint.sh /bin/entrypoint.sh

ENTRYPOINT ["/bin/entrypoint.sh"]
