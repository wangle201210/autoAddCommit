FROM alpine

ENV APP_ROOT /var/www
ENV APP_PATH $APP_ROOT/test-one

RUN mkdir -p $APP_PATH

ADD ./main $APP_PATH/
ADD ./*.sh /bin/
RUN chmod +x /bin/*.sh


# 运行命令
CMD ["make.sh"]
