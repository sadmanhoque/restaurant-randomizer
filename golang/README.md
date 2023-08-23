### Introduction ###

This is meant to implement the restaurant-randomizer project but using Golang for the backend, I haven't yet decided about the front-end, might to React, or experiment with a Golang front-end server, could to both honestly.

## Project setup ##

For Golang setup in Windows go to https://go.dev and install using the msi file, for Mac or Linux it would likely recommend using a package manager like brew or apt.

For Windows at least, you could run into an error about gopls not finding the module, this happens because by default golang expects you to run all projects out of a specific directory, you can find out what it using the command: `go eng GOPATH`. I find this an annoying limitation since I like to keep all my projects in a git directory, so to get around this, if you want, run this command while in the same directory as the main go file: `go mod init main`.