FROM quay.io/operator-framework/upstream-opm-builder AS builder

FROM centos


ARG USER_ID=1001
ARG GROUP_ID=12009


RUN groupadd -g ${GROUP_ID} myuser \
 && useradd -g myuser -u ${USER_ID} -m myuser \
    usermod -aG wheel myuser

#ENV USER opmuser
#RUN groupadd -g 12009 ${USER} &&\
#    useradd -g ${USER} -u 1001 -m ${USER} &&\
#    usermod -aG wheel ${USER}

#ENV USER opmuser
#RUN groupadd -g ${GROUP_ID} ${USER} &&\
#    useradd -g ${USER} -u ${USER_ID} -m ${USER} &&\
#    usermod -aG wheel ${USER}


USER ${USER}
RUN whoami

LABEL operators.operatorframework.io.index.database.v1=/work/index.db

COPY --chown=myuser:mygroup ["nsswitch.conf", "/etc/nsswitch.conf"]
COPY --chown=myuser:mygroup ["database", "/work"]
COPY --chown=myuser:mygroup  --from=builder /bin/opm /bin/opm
COPY --chown=myuser:mygroup --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

RUN chown  myuser:mygroup /work/index.db &&\
    chown  myuser:mygroup /bin/opm &&\
    chown  myuser:mygroup /bin/grpc_health_probe &&\
    chown  myuser:mygroup /etc/nsswitch.conf
  

EXPOSE 50051
USER 1001
WORKDIR /work

ENTRYPOINT ["/bin/opm"]
CMD ["registry", "serve", "--database", "index.db"]
