///func main, func client, func server

package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "time"
        "flag"
)

/* 
Rules: 
        1. best 2 out of 3
        2. Your code should be able to act as either a client or server.
        3. your rock-paper-scissors program should be able to accept input from the command line, so that a human can play 
           interactively, and it should also be able to play automatically, submitting moves according to some strategy you
           have coded up. 
        4. efficiently handle errors 

Tips:
        1. use your knowledge of socket programming to complete the task
        2. work out a common protocol (set of rules for the format and order of the messages that are exchanged between the two sides)
                - must agree upon things such as which of the three rounds you are on, who made which moves, who won, etc.
 */


func main() {
        playerType := flag.String("player", "Beginning game now...", "Are you a computer or a human?")
        chooseOpponent := flag.String("opponent", "Beginning game...", "Are you playing a computer or a human?")
        
//         ipAddress := flag.String("ipAddress", "169.229.50.178", "INPUT IP ADDRESS")
//         port := flag.Int("port", 2003, "INPUT PORT NUMBER")
//         flag.Parse()

        if *chooseOpponent != "" and *playerType != "" {
                if *playerType == "human" {
                        if *chooseOpponent == "computer" {
                                fmt.Println("Beginning game...")
                                clientcomp(*ipAddress, *port)   //if a human chooses to play the computer, then the client code will run
                        } 
                        // else if opponent is a human!! in this case, you have to be both a client and a server! so have two more functions
                } else if *playerType == "computer" {
                        if *chooseOpponent == "human" {
                                fmt.Println("Waiting for human player...")
                                servercomp(*port)   //computer will act as a server, and respond to the client human's moves
                        }
                } else { 
                        fmt.Println("Please enter who you are, so that the game can begin.")
        } else {
                fmt.Println("Please enter your player type (human, computer) and opponent type (human, computer)!!")
        }
}

func clientcomp(ipAddress string, port int) {
        clientConn, err := net.Dial("tcp", IPAddressPort)
        if err != nil {
                fmt.Println("ClientComp Connection Error:”, err)
                return
        }
                            
        reader := bufio.NewReader(clientConn)
        numGames := 3
        
        for i := 0; i < numIters; i++ {
        
               //insert strategy for playing, and depending on whatever the opponent plays, initialize "sendMsg" to what could beat that!!
                //if blah blah:
                //    sendMsg := blah blah rps
                // elif blah blah:
                //    sendMsg := blah blah rps
                // then, continue on by printing and sending that message!!!
        }                    
                            
        clientConn.Close()

}

func servercomp(port int) {

        serverConn.Close()
}



///will need to create a separate client function to connect with John's client function!!
