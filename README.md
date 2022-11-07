## Project: 
Net-cat

## Authors: 
* Jacob Pesämaa
* Nafisah Rantasalmi
* Huong Le

## Technologies:
Project is built with Golang version 1.18

## Description: 
The project recreated the NetCat in a Server-Client Architecture that can run in a server mode on a specified port (:8989) listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server. Maximum 10 connections.

## Basic Features:
* The TCP Chat, if not otherwise specified, will be running on a default port (:8989)
* The TCP Chat receives up to 10 clients
* These clients will be able to communicate with each other, receiving notifications when a client joins or leaves the chat
* A new client will be able to see all the previous messages

## Set up: 
1. Clone the repo in VSCode or your text editor of choice.
2. Go build the program: 
```
$ go build -o ./TCPChat
```
3. Run the program: 
```
$ ./TCPChat
```
4. To run the program on a different port:
```
$ ./TCPChat <port>
```
5. To end the program, type 'x' to the terminal

## Usage:
* From the same computer:
1. Type the command into a new terminal (client's) to start:
```
$ nc localhost <port>
```
2. When connection is received, a linux logo would appear and ask for client's name
3. Enter your name and start typing!
4. For more than 1 client, open new terminals (maximum 10 connections), then use the same command and port to join the chat.
* From different computers (Mac OS): 
1. Find host IP Adress by using command: 
```
$ ifconfig
```
2. Look for CSUM> inet (address)
3. On second, third, ... computers; type into terminal command:
```
$ nc <address> <port>
```
3. Start chatting!

## Implimentation Details:
// Empty history log with func EmptyHistory

// Port validity check:
- func ValidPort: Port number is only numeral and is valid from 1024 to 65535
- func CheckPort: If port is 8989 or from 1024 to 65535. The port is valid. Otherwise return "[USAGE]: ./TCPChat $port"

// Listen with a valid port:
- The port is now valid and running
- func Exit: Print "[Type 'x' to quit]" - Type 'x' to exit, if input is different than 'x', func Exit will be recalled
- Listening on port: ":" + [port]
- If there is error (func CheckError), return "ERROR: unable to listen to port" and closes the listener

// Accept and handle up to 10 connections with go routine:
- Accept connection, if err then exit
- Once the connection is accepted, a welcome message will be printed and asks for client's name ("welcome.txt" + func GetName)
- If a new client joins the chat, and if the log ("history.txt") is not empty, it will also print out the log
- If the number of connections (var int) is more than 10, client will receive the notification "Chat is full"
- Go routine: HandleClient: to send join/leave notification to all clients via channel and to send messages to all clients via channel under the format: [Time] (func CurrentTime) [Name] [Message]
- When a message is sent, func AddToFile will write the string into the log 
- Go routine: Broadcast: Combine the broadcast into one function
- Once a client leaves the chat, delete that client, name and close connection.

// Struct models, map of client and channels can be found under 'server.go'
