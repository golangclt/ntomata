# ntomata

A CLI time management tool based off working in focused, time-boxed lengths of time, then taking breaks. Basically, the pomodoro technique.

ntomata is greek for tomato. The project was called pomodoro, but this was cooler.

If you are on OSX -- it will notify you when time is up through a terminal notification and a nice voice message. 

## Usage
`start`: start the ntomata. (default 25m)  
`sbreak`: take a short break (default 5m)  
`lbreak`: takea long break (default 10m)  

## Why? 

Created as a demo app to show off some of go's features as part of a talk.

## Develop 

This project only has one external dependency, `deckarep/gosx-notifier`  

To develop, simply `go get github.com/deckarep/gosx-notifier` and use the standard go tool chain.

Let's improve this codebase as a group! There are many ways to improve it.

## Todo

- A cool notification for Windows and Linux users.