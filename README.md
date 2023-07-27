# SeaSorter

![https://iili.io/HQyYXSa.png](https://iili.io/HQyYXSa.png)

Simple desktop app for sorting files by their formats, written with [fyne](https://iili.io/HQyYXSa.png) framework.

## Install

**Windows**

Quick way to get SeaSorter for windows id to [download](https://github.com/ustymhentosh/SeaSorter/blob/main/SeaSorter.exe) `Seasorter.exe` from this repo.

**Linux**

Linux user can also just [download](https://github.com/ustymhentosh/SeaSorter/blob/main/SeaSorter) `SeaSorter` file and use it as is. If you don't have specific app to open executables → try open `Properties` of the file, and tick `Executable as Program` badge.

![linux_4.png](https://github.com/ustymhentosh/SeaSorter/blob/main/images/linux_4.png)

For further distribution you can download whole [fyne-cross](https://github.com/ustymhentosh/SeaSorter/tree/main/main/fyne-cross) folder in main directory of this repo.

## Usage

Application should look like this.

Basically, you just enter format of the file and folder where it should go, if folder does not exist, it will create one, if exists it will just move file there, you can move multiple files into one folder.

![usage.mp4](https://github.com/ustymhentosh/SeaSorter/blob/main/images/usage.mp4)

You can also use automatic sorting, which just creates separate folder for each format.

![usage-auto.mp4](https://github.com/ustymhentosh/SeaSorter/blob/main/images/usage-auto.mp4)

Application does not work with files in child folders of the main folders.

Other functionality include date filtering - specify how old file should be to be affected by application.
