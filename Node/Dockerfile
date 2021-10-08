FROM node:14    

WORKDIR /usr/src/app

COPY package\*.json ./

RUN npm install

RUN npm i nodemon -g

COPY . .

EXPOSE 3500

CMD ["nodemon", "index.js"]