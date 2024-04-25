FROM node:21-bookworm-slim

WORKDIR /web

COPY package.json /web/package.json

RUN npm install -g pnpm@9.0.4

RUN pnpm install