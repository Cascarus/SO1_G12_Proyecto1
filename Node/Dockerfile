FROM node:14    

WORKDIR /usr/src/app

COPY package\*.json ./

RUN npm i nodemon -g

COPY . .

RUN npm install

EXPOSE 3500

CMD ["nodemon", "index.js"]