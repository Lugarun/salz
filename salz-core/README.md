# Quickstart
Getting the docker container up and running takes a long time and right now needs
to redo itself everytime there is a change in a file.
## Via Docker
Run `docker build .` within this folder to just build salz-engine.
Run `docker-compose up` to build and run this in conjunction with the rest of the project.
## Via Stack
In order to run the game engine you need to run the database (it is how you interface
with the game engine).
To do this you build and run the docker file in the `salz-db` folder (there are
scripts in that folder to build, run, and connect to the database).
You then build this project using `stack build` in this folder.
Finally you can start the game engine and connect it to the database in the ghci which you
start with `stack ghci` by running the command the command in the `./app/Main.hs` file.
