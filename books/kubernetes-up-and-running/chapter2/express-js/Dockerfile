#Star with a Node.js 10 TLS image
FROM node:10

# Especifie a image dir in witch all comandas will be executedes
WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install 

COPY . .

CMD [ "npm","start" ]
