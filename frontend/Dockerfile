FROM node:20-bookworm-slim

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN pnpm install

ENV PORT=8080 \
    WATCHPACK_POLLING=true \
    NEXT_TELEMETRY_DISABLED=1 \
    NODE_ENV=development

CMD [ "pnpm", "run", "dev" ]