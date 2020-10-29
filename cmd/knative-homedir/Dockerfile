FROM registry.access.redhat.com/ubi8/ubi-minimal
COPY ./build/_output/bin/knative-homedir /knative-homedir
WORKDIR /
CMD ["/knative-homedir"]
