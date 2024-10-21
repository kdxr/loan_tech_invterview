FROM nginx
#ARG user=swadmin
# ENV HOME /home/$user
# RUN adduser -D $user
# RUN mkdir -p /usr/src/app
COPY ./default-http.conf /etc/nginx/conf.d/default.conf
# WORKDIR /usr/src/app
# RUN chown -R $user:$user /usr/src/app
USER $user
EXPOSE 80
