

## Filely

## Overview
Filely is a web app that sends files to a server using Transfer Communication Protocol. It demonstrates  the interaction that occurs between the  client and server.files are uploaded and sent from the client's side. If the server is on, it accepts and saves the files. it has  a user-friendly interface.

## preview

![filely-screenshot][client/public/filelyimage.png]




## Keep in mind
- filely contains the client and server side running on different ports
- client side is running on localhost/8080
- server side is running on localhost/9000


## Installation:
- Clone the repository


- cd into the clone repository

- Open two terminal seperately for the client and server.

### for the server side
- cd into the server folder 
- run server.go in the terminal
-  you should see TCP server started on port 9000 in your terminal
- go into your browser and enter localhost/9000
- check your terminal to see Client connected from: [::1]:62906(this numbers are not fixed)

### for the client side
- cd into the client folder 
- run main.go in the terminal
-  you should see listening on port 8080 in your terminal
- go into your browser and enter localhost/8080



## Usage
- Choose a file from the client side
- Click the send button after the file has uploaded
- If its successful a confirmation message is sent to the client side
- The server stores the file in the receive.txt file
- The file is uploaded in your project 

## Technologies Used:
- Go
- HTML
- CSS


 ## Contributing

 If you want to contribute or have suggestions, feel free to open an issue or submit a pull request.


 ## Resources

  https://www.youtube.com/watch?v=DHlV65WTG3k (styling file input button)

  https://freshman.tech/file-upload-golang/ 

