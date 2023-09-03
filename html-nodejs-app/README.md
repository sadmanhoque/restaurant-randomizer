### restaurant-randomizer ###

Steps to get the webApp working:

* Get your own Google Maps API key and dump it in a .env file in root, looks something like this: 

`GOOGLE_MAPS_API_KEY="askdjfbnsrkjfn"`

Make sure that this API key has the correct permissions from the GCP console, the free tier limit is good enough, I believe its specifically using the Places API, which must be enabled for the issued API key. Also make sure that GCP limits the API usage to the free tier if you don't want unexpected bills, it said it's setup that way by default, but mine was not. 

* The command to get the server running is 'npm run dev' and the default url is localhost:8383.