FROM node:21-bookworm-slim

WORKDIR /web

COPY package.json /web/package.json

RUN npm install -g pnpm

RUN pnpm install