### DGF - Dumb Game Format

DGF is a simplified and fractured version of SGF. For the sake of prototyping an also creating an extreamly lightweigth platform for playing go, DGF seeks to remove the exelent anotation storing features of SGF, while adding a more rigid structure to allow more ineroperability between the human redable file and computer's memory representation.

_NOTE_: DGF's goal of being more strict than SGF with regards to file structure comes from a desire to be able to act as a text-based or binary file with optimized space usage. The firt implementation ignores this by using configuration file storage for game files 


DGF is primarily useful as a standard interface between software using SGF. SGF itself is standard across the board (normally) but the way it is internally represended in memory is absolutely not. By having an intermediary file formart with a simpler, numerfical view of boards states, individual applicaitons can rely on entirely different sgf parsing libraries by simply writing one or two functions that translate a DGF move to the appropriate input for the SGF parser, and that reads an SGF Parsedrs board structure to generate a board state.

## Feature tracker
 * Standard
  * game
  * Board State
  * Move
 * Functionality
    * Move from board state diffrence
    * Move validation
        * KO
        * Self-capture


#### DGF macro structure

DGF's utility is to think of a game's representation the way a computer would, making it numerically efficient to parse and validate boards and moves. However an additional idea being played with is the ability to extend DGF into meta data and perhapse even SGF's realm of annotations.

DGF has three units: a game file, a board file, and a move file. The purpose of this is to compartimentalize the users actions. It creates intuitve interopperation between extreamly small components.

The modular notion also allows needed complexity to be selected by the using applicaiton.

#### DGF Binary

DGF has a standard binary format designed to make it clean for applications to store data in a strict way without resorting to marshalled program variabels. __This is not implemented yet.__

##### Game file

A game file a contains everything relevant to a game, including latest board state and complete move sequence.

Game file contains the follwoing fields:
Game ID: a unique identified for the application to use in cataloguing and adding further ingormation
Black: Name or ID of the player playing black
White:
Date: Date the game started
Initial board state: The initial board state array. Contains information regarding komi + handicap
sequenc eof moves: array of move files played

##### Move file

A move is defined as a single int, represeting the position on the board the player wishes to place a stone. 1 indexed.

A move does not specify a color and _must_ reference a board state.

i.e. the 3-3 point at SGF "dd" would be 80.

A passing move is 0. Resignation is a negative value.

#### Board/Position file

A go board is represented by a single array of bytes. For an n x n board, the array is of length n^2 + 3, the first three bytes being equal to :
which player's turn it is (0 = game is finished, 1 = black, 2 = white)

Byte 1
Magic byte:

bit positions:

8765 4321

bytes 1 & 2 are for player turn

1 -> black 2 -> white

if 2 && 1 == flase -> game is over

bits 3 & 4 are for komi half points (3 -> black, 4 -> white)

bits 5 & 6 are for capture overflow
bit 5 -> black has + 255 captures, bit 6 1 -> white has +255 captures

Byte 2
Black's captured stones + reverse komi

Byte 3
White's captured stones + komi

Byte 4
board size

Following that each byte is the state of the intersection going left to rigth, top to bottom.
0 = empty
1 = black stone
2 = white stone

For evaluations, it can be useful to note that nodes 1, n, n^2 - (n-1) and n^2 are corners, nodes 1-n, n^2 - (n-2) to n^2 - 1 as well as any node fro which idex % n = 0 or 1 AND idex > n * (n - 1) are sides. As the game of go needs only count libertides, this can be generalized relatively easily, but needs to be checked for as lots of natural adjacency operations would lead to wraparound. 

TEMP; adjacent coord func (i, is the 1 indexed position, n is the size of board)

```if i == 1 // 1-1 conrner
if i == n // 19-1 corner(or equiv)
if i == n\*(n-1) + 1
if (i + 1) % n == 0 // Left side, 1-x
if i % n == 0 // rigth side n-x
if i < n // top side x-1
if i > n \* (n-1) // bottom side

return i-1, i+1, i-n, i+n 
```
Fun note: so long as intersectoins can only have 4 states (i.e. empty, blakc, white, illigal ko or other), pairs of 4 stones can each be represented by a single 8-bit byte. A 19x19 boardstate could fit in as little as 70 bytes!


### Record file

A record is a representation of the whole game. Currently this excludes time controls, which as still treated as metadata.

Memory wise, a recod is an array of integers, The first 4 + n^2 integers are the bytes of the original boardstate, the following integers is the result byte (0 = unfinishes, 1 black victory, 2 white victory) following which are all the moves that were played.

### Board state update

Given a move and a board state, generate the new board state.

### Move validations

Given a board state and a move, determining if that move is legal. The following contidiions are checked for:

#### Self-capture

### Move from board state diffrence

Given 2 board states, return the move which created the new board.

