# Number Guessing Game
- Simple number guessing game where the computer randomly selects a number and the user has to guess it. The user will be given a limited number of chances to guess the number. If the user guesses the number correctly, the game will end, and the user will win. Otherwise, the game will continue until the user runs out of chances.
- This project is from the project based road map, you can find it [here](https://roadmap.sh/backend/projects).
- You can read more about the project form [here](https://roadmap.sh/projects/number-guessing-game).

### About
- When the game starts, it displays a welcome message along with the rules of the game.
- Users should select the difficulty level (easy, medium, hard) which will determine the number of chances they get to guess the number.
- If the user’s guess is correct, the game should display a congratulatory message along with the number of attempts it took to guess the number.
- If the user’s guess is incorrect, the game should display a message indicating whether the number is greater or less than the user’s guess.
- The game ends when the user guesses the correct number or runs out of chances.

### Features
- Timer to see how long it takes the user to guess the number.
- Keep track of the user’s high score (i.e., the fewest number of attempts it took to guess the number under a specific difficulty level and Time the user took.).
- All configuration can be found in a easy to read json file.

### Usage
```
./NGG

Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Choose one of the following:
1.Start Game
2.Exit
Enter Your Choice: 1

Please select the difficulty level:
1.Easy (10 chances)
2.Medium (7 chances)
3.Hard (4 chances)
4.Back
Enter Your Choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess (7 Chances, ELapsedTime 3.6e-07): 55
Incorrect! The number is less than 55.

Enter your guess (6 Chances, ELapsedTime 1.035469544):
```
