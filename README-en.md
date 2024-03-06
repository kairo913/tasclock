[README in Japanese](https://github.com/kairo913/tasclock/blob/main/README.md)

# TasClock: Task Management and Time Tracking App

TasClock is an app that combines task management and time tracking. It measures the time spent on each task and calculates the current hourly rate based on the set unit price.

## Main features

-   Create, edit, and delete tasks
-   Record start and end times for each task
-   Calculate the current hourly rate in real time based on the set unit price
-   Display a graph of the work time history for each task
-   Set a target hourly rate to stay motivated
-   Export data in CSV format

## Dependency

Go 1.18+  
NPM (Node 15+)

## Install Wails

`go install github.com/wailsapp/wails/v2/cmd/wails@latest`

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.  
This will compile your project and save the production-ready binary in the `build/bin` directory.

## Reference

[Wails: Installation](https://wails.io/ja/docs/gettingstarted/installation)

[Wails: Compiling your Project](https://wails.io/ja/docs/gettingstarted/building)
