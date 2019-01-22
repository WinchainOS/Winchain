FROM xwc1125/go:latest
MAINTAINER xwc1125

ENV DATA_HOME /data
# 添加项目
ADD . /Winchain
RUN cd /Winchain && make all
RUN mkdir -p $DATA_HOME/Winchain

# 配置环境
RUN cp /Winchain/build/bin/geth /usr/local/bin/
RUN cp /Winchain/build/bin/bootnode /usr/local/bin/

RUN mkdir -p /example
RUN cp -rf /Winchain/docker/data/* /example/
RUN cp -rf /Winchain/docker/sh/* /bin/
RUN chmod +x /bin/*

RUN sh /bin/init.sh

# 删除源码
RUN rm -rf /Winchain

VOLUME /data
EXPOSE 9545 3000 30303 30303/udp

ENTRYPOINT ["/bin/entrypoint.sh"]
CMD ["start", "-D"]
