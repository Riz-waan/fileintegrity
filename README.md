# README

## About

This is a sample project done as a proof of concept for a file integrity verifier based on the CIA Triad.

Current Features

* Ability to add single files
* Ability to remove files
* Ability to see between what timestamps the file was corrupted
* Files will be checked every 30 seconds

Future Features

* Clean up the UI
* Allow user to set File Checking Interval
* Allow user to add multiple files
* Allow user to add folders with recursive checking for the files
* Allow the application to run in background as taskbar item
* Have notifications when file is compromised
* Allow the sorting of the table and searching

## Running and Building

This project is created in the wails and you must install wails before running it [Installation | Wails](https://wails.io/docs/gettingstarted/installation)

To run in live development mode, run `wails dev` in the project directory.

To build a redistributable, production mode package, use `wails build`.
