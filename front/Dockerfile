FROM node:lts-alpine
WORKDIR /var/bbs/front/
COPY ./package.json .
# COPY ./package-lock.json .
COPY ./babel.config.json .
COPY ./webpack.config.js .
COPY ./index.html .
RUN npm i

EXPOSE 8080
CMD ["npm", "run", "start"]