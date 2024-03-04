# Go Debug Visual Studio Code

## Title

Swiss knife for Go debugging with VSCode

## Small Print

Being able to debug your code in the IDE should be an easy process. If you ever struggled with debugging Go code in Visual Studio Code, this session is definitely for you. We'll take a look at how to debug several kinds of projects by only tweaking settings in the dreaded "launch.json" file. You'll discover how many options and customizations you can apply to this file to leverage your debugging experience. Finally, we briefly touch on the key features of the Delve debugger.

## Description

One of the most important things when we're writing code is the ability to debug it. Many IDEs have an integrated debugger that can smoothen our coding experience. The debugger for the Go source code is called Delve. It's tightly integrated with VSCode and the Go extension. As you might know, the debugger allows us to step through our code, focus on specific sections that may deserve more attention, inspect variables' values, stack traces, etc.
Sometimes, debugging turns into a hassle. The process supposed to help us becomes an insurmountable obstacle. Sometimes, we abandon the debugging or log directly into the code. Both of the options end up decreasing our productivity as developers.  

Thus, this talk aims to provide a working solution to debug Go code in VSCode. I chose this IDE since it's free, highly customizable, performant, and my favorite!
Since we can build different projects, I try to provide you with a working solution for each. The scenarios you're likely to face are (list not exhaustive):

- Debug unit tests
- Debug integration tests
- Debug a package
- Attach to an already running process (both locally and remotely)
- Debug multiple microservices with the compound configuration

To overcome these challenges, you're requested to tweak settings in the `launch.json` file within the hidden .vscode folder. In this file are listed what are known as profiles. These are selectable in the "Run View" area. Within this file, we have the option to set different values, such as:

- environment variables or env file
- whether to show global variables' values
- which console to use
- and many others

Throughout this talk, I share hints on Delve and how to make the most of it. I also touch on some overlooked aspects of VSCode that can make a huge difference in your debugging experience.
Finally, I'll give you some mind-blowing tips and tricks on debugging. If you'd like to improve your debugging skills, please don't miss my session!

## Outline

The outline should be as follows:

- Intro: 3 minutes
- Why Debugging: 3 minutes
- Delve Debugger: 3 minutes
- Debugging in VSCode: 5 minutes
- The launch.json file: 5 minutes
- launch.json's options: 5 minutes
- Outro: 1 minute
