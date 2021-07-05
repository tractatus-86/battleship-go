# Battleship-go

## How to run
To get started run 
```sh
docker build -t battleship .
docker run -it battleship
```

```sh
Usage of ./battleship:
  -input_file string
        input file of commands
  -output_file string
        output file of responses
  -pretty
        printer mode
```

Inside the docker container run to get a example game
```sh
./battleship -input_file fullgame_test -output_file output
```
or
```sh
./battleship < fullgame_test > output
```

or for interactive mode

```sh
./battleship -pretty
```
or for animated output mode
```sh
./battleship -input_file fullgame_test -pretty
```

## How to play

*Note : The command interpreter is case insensitive*

Grid
```
    1   2   3   4   5   6   7   8   9   10   
A | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
B | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
C | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~
D | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
E | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
F | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
G | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
H | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
I | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
J | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ | ~ 
```

Ships 
```
Destroyer
Carrier
BattleSHip
submarine
cruiser
```

### Phases
There are thee phases Setup, Battle, Gameover

#### Setup
Place all ships
```
PLACE_SHIP Destroyer right A1
```

#### Battle
Fire until all ships are destroyed
```
Fire A1
```
#### Gameover
All ships are destroyed