# Dvd Game Launcher
Dvd Game Launcher is a simple launcher for video games, letting you put together a disc with multiple DRM-free and titles along with any extras, such as manuals, artwork or OSTs and present them through one simple menu.

## Why
With the preservation of digital media being more important now than ever, Dvd Game Launcher is an artistic approach to maintaining those efforts. Inspired by the brilliant program, DvdStyler, it allows users a creative way to simply accessing their games, without having to actually navigate the dvd's file structure when deciding to play from it.

Because it is lightweight and only depends on being present within the files, backing up the dvd to a drive or computer won't have any impact.

## Stack
Dvd Game Launcher is written entirely in Go. It utilizes the Fyne package for its UI elements.

## How To
### Expected DVD File Structure
Assuming control over the internal file structure for your DVD is not within the scope of this project, it is important that the dvd file structure looks like this:

``` text or
/Collection
    launcher.exe (this program)
    /Game1
        <your game exe here>
        <other game files here>
    /Game2
        <your game exe here>
        <other game files here>
    /Extras
        Manuals/
        OST/
        Artwork/
    /assets (explained below)
        ui_background.png
        logo.png
```
You may also check the examples folder in this repo where examples/structure is your game collection. 

The only file changes this program will make is to output save data to a local drive, as saves will not be written to the dvd.

### Asset Structure

This repository includes an assets/ folder containing development placeholders used by the launcher during testing.

When creating a DVD, you must separate /assets folder next to launcher.exe containing the runtime assets the launcher will load. To maintain user flexibility, it is set up to where the user can use their own image to personalize their setup.

