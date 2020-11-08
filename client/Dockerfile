FROM node:alpine as build
WORKDIR /app

COPY package.json ./
RUN npm install
#RUN npm ci --silent
#RUN npm install react-scripts@3.4.1 -g --silent
COPY . .
RUN npm run build

# production environment
FROM nginx:stable-alpine
#COPY proxy.conf /etc/nginx/conf.d/default.conf
#COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
