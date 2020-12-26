# Poker using Go

![Gopher](/GOPHER_PARAKEET.png) 

**Overview**

This is an application to manage the win count of players.

Go has ease of extending across multiple delivery mechanisms and storage components.

This project uses an in memory store for performing the tests quickly and file system store to persist data when used as an application.
The application can also easily be scaled to use a database like Postgres but implementing the PlayerStore interface available.

**The application is delivered using various delivery mechanisms.**

1. REST APIs.
    * GET /players/{playerName} : Returns the win count of the player in text format.
    * POST /players/{playerName} : Increments the win count of the player by 1.
    * GET /league : Returns the list of players along with their win count in Json format.
2. Command Line Interface
    * Prompts the user to enter the number of players to begin the game.
    * Notifies the current blind value being altered by a scheduler.
    * Game can be closed by entering the winner's name followed by "wins".
3. Web Sockets
    * GET /game opens a web page with a text box to enter the winner's name.
    * The websocket connection records the winner's name and includes the information to the persisted data.
   
**Testing Criteria**
 
1. Spies and Stubs are used to test the functionality without the tests being concerned about the dependencies. 
2. Test helpers have been written to making the tests look cleaner and sensible.
3. Tests have been grouped into relevant blocks making them more organized.
4. Slices have been used to run multiple tests with similar arguments of different values, making the tests parameterized.