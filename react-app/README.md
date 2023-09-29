### Intro ###

React+Vite project utilizing the tailwind CSS follow the setup instructions here https://tailwindcss.com/docs/guides/vite

To run server use the command `npm run dev`

FYI this is only the instructions for the frontend component of the application, this is setup for the backend-golang project
in this same repo. To successfully run the application both servers need to be running. For instructions on the backend server use the Readme file in its directory.

### Containerization ###

The Dockerfile and changes to vite.config.js file are already setup as required (default port is 5173), we only need to make the image and run it.

Note that the app in this directory and hence the docker image only contains the front-end react app, not the backend, the website won't actually be able to do anything without the backend app running.

Sample commands:

`docker image build -t react-app:dev .`

`docker run -d --rm -p 5173:5173 --name react-container react-app`

## TODO ##

-setup environment var for port, backend server URL, backend server port.
-update the API and UI to also utilize the location, perhaps using a map? Checkout 'leaflet'.
-update the API and UI to utilize the search range, this could also be map based, maybe.