FROM node:lts-alpine

# install simple http server for serving static content
RUN npm install -g http-server

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY halalfind-frontend/package*.json ./

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY halalfind-frontend/. .

# install project dependencies
RUN npm install

# build app for production with minification
RUN npm run build

EXPOSE 3000

CMD ["npm", "run", "dev"]

#CMD [ "http-server", "dist" , "--port=3000"]