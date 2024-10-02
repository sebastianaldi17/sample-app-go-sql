FROM node:22-alpine

RUN mkdir -p /ts-server

WORKDIR /ts-server

COPY package.json ./
COPY tsconfig.json ./
COPY src ./

RUN npm install && npm cache clean --force

EXPOSE 3000

CMD ["npm", "run", "dev"]