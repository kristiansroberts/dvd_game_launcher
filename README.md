# Dvd Game Launcher

## How To
### Expected DVD File Structure

It is important that the dvd file structure looks like this:
``` text or
/Collection
    launcher.exe (this program)
    /Game1
        <your game here>
        <other game files here>
    /Game2
        <your game here>
        <other game files here>
    /Extras
        Manuals/
        OST/
        Artwork/
    /assets (explained below)
        ui_background.png
        logo.png
```
You may also check the examples folder in this repo where examples/structure is your game collection

### Asset Structure

This repository includes an assets/ folder containing development placeholders used by the launcher during testing.

When creating a DVD, you must separate /assets folder next to launcher.exe containing the runtime assets the launcher will load. To maintain user flexibility, it is set up to where the user can use their own image to personalize their setup.

