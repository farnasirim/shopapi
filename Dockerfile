FROM quay.io/brianredbeard/corebox

ADD shopapi /bin/shopapi

ENTRYPOINT ["/bin/shopapi"]
