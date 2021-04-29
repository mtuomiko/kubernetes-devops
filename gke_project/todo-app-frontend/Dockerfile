FROM node:14-alpine as build
WORKDIR /work
COPY package.json ./
COPY package-lock.json ./
RUN npm install

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /work/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]