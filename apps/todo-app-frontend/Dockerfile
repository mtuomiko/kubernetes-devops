FROM node:14-alpine as build
WORKDIR /work
COPY package.json ./
COPY package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /work/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/nginx.conf
RUN rm /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]