# Blue Archive Wiki Strategy Maker 3

## Usage
* Run `build.sh` to build the parser.
* Change Directory to `bot/` using `cd bot/`.
* Create a file called `.env` in `bot/` and fill the information like so:
```bash
BOT_USER="username@botname"
BOT_PASSWORD="botpassword"
BOT_CONTACT="At Blue Archive Wiki Discord @username"
# Or
# BOT_CONTACT="At Blue Archive Wiki User_talk:username"
```
* Run `login.sh` and `get_csrf_token.sh`.
* Run `bot_edit.sh`. This will fetch the wiki pages and update the pages locally using data from `strategies/input/`.
    * While `bot_edit.sh` is running, you will get prompted with a diff of the changes, and after leaving the `less` command asked for approval.
* Run `bot_submit.sh` when ready to submit the edits.

## Making Strategies
### File Naming 
First thing is naming the file.<br> Here are the requirements:
```
${MissionCode}.${TabOrderInt}.${StrategyTitle}.txt

# Normal Missions
6-1N.1.3-Turn Clear.txt
6-1N.2.3-Turn Alternative Clear.txt
6-1N.3.240s Time Challenge.txt

# Hard Missions
10-2H.1.4-Turn Challenge.txt
10-2H.2.6-Turn Gift Clear.txt
```
There are prenty of edge cases of which strategy should go first in Tab Order, but basically it boils down to:

1. 3-Star Clear.
2. Challenge.
3. Gift Clear.
4. X Only.
5. Alternative Clear.
6. Time Challenge.

You can read this as a Priority List, as in, if you can do 1, 2, and 3, that gets higher priority, higher as in being at the top, than just being able to do 1, so if you do 1, 3 and 1, 2 in different strategies you can put 1, 2 first then 1, 3 later.

Point 4 is to use if you can't get 3-Star plus challenge or gift clear, so it becomes Challenge Only, or Gift Only.

You can use, more or less, the same list for strategy title, but add the number of turns, or do `X Seconds Clear` Time Challenge with X being the number of seconds.<br>
Look at `strategies/input/` for examples.

### Writing Strategies
Writing is the similar to how you would read a startegy, but with less words, and more robotic.
```
team ...
turn
...
turnenemy
...
```
#### Steps:
1. Add the teams 
2. Type `turn`
3. Type an `Action` for each team
4. Type `turnenemy`, or go to step 2
5. Do step 3
6. Repeat steps 2-5 until you first use the boss `Action`

#### Actions
This lists all actions the program can handle, and what to type for each action.
```
bawsm3 -actions
```
Rule of thumb is: `Action Letter Direction Armor Extras`, But some actions don't need certain arguments, they might need two of the same, or require some exclusive arguments.

#### Tips And Recommentations
##### General
* If the team can withdraw from a fight and still get 3-Star use `Withdraw` Action, but that retreats the team btw.
* When routing time challenges 3-Star is not a concern.
* If Possible try to route for 1 team to have 1 armor type to deal with.
* You can write multiple strategies, but one, or more if necessary, for 3-Stars, Challenge, Gift, and another one separate for Time Challenge is all you need.
* No need to say UpRight/DownLeft on a direction when you can only go to one kind of up/down.
* Do Not Say Right/Left when you mean to go UpRight/UpLeft or DownRight/DownLeft.
* Run you Strategy to tho the program to see if it doesn't have simple errors.
* Up/Down is used for direction and Top/Bottom is used for position. (You can also use Middle/Center for position)

##### Situation Awareness
Some times an action won't trigger what, or happen when you think in-game.<br>
Examples:
* `Swap` Won't trigger a teleport popup confirmation.
* `StartTeleport` Caused by a Broken Tile doesn't Happen on Enemy Turn it happens at start of Player Turn.
* `StartTeleport` Caused by a Player Action Will happen right after said Action.
* `Attack` If an Enemy occupies the same Tile as a Gift you can still obtain the gift, so type Attack with Gift as an extra argument.
* `Move` Order Is important if Team A moves closer to an enemy that is going to attack on his turn he will give priority to the one that moved closer first.

##### Action Ordering
* Try to make the teams go in alphabetical order. (This gives more importance to the action when the order is broken)
* The Team that got closer to an Alerted Enemy first is the one that is gonna be attacked.
* Team A should be the one killing the boss.
* Declare the teams in alphabetical order.
