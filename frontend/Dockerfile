FROM node:20 AS builder

WORKDIR /app

COPY ./frontend/package*.json ./

COPY ./frontend .

RUN rm -rf node_modules package-lock.json
RUN npm install
RUN chmod +x ./node_modules/.bin/*
RUN npm run build 

FROM nginx:stable-alpine AS final

RUN rm /usr/share/nginx/html/index.html

COPY --from=builder /app/dist /usr/share/nginx/html

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]