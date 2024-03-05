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

## How to build

### Dependency

Go 1.18+  
NPM (Node 15+)

### Install Wails

`go install github.com/wailsapp/wails/v2/cmd/wails@latest`

### Build

`wails build`