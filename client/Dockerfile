FROM node:18-alpine AS builder

WORKDIR /app

COPY ./app/package*.json ./
RUN npm ci --ignore-scripts

COPY ./app .
RUN npm run build

FROM nginx:stable-alpine AS runner

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 5173

RUN mkdir -p /var/cache/nginx && chown -R nonroot:nonroot /var/cache/nginx 
RUN mkdir -p /tmp && chown -R nonroot:nonroot /tmp
RUN mkdir -p /run && chown -R nonroot:nonroot /run

USER nonroot

CMD ["nginx", "-g", "daemon off;"]