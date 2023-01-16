FROM node:lts as js-builder
WORKDIR /web
COPY web/package*.json web/tsconfig*.json ./
ENV PATH /web/node_modules/.bin:$PATH
ADD web/public/ public/
ADD web/src/ src/
RUN npm install
RUN npm run build
CMD npm start
