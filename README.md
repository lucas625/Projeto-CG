# Computer Graphics Project

## Table of Contents
- [Computer Graphics Project](#computer-graphics-project)
  - [Table of Contents](#table-of-contents)
  - [Description](#description)
  - [Requirements](#requirements)
  - [Code Extensions](#code-extensions)
    - [VS Code](#vs-code)
  - [Installing](#installing)
    - [Cloning](#cloning)
    - [Golang](#golang)
    - [Dependencies](#dependencies)
  - [Input](#input)
  - [Running the Project](#running-the-project)

## Description

A Path Tracing project developed for computer graphics discipline at Federal University of Pernambuco.

## Requirements

- go 1.13
- gomod
- git

## Code Extensions

### VS Code

- Markdown All in One
- GitLens
- GO

## Installing

Follow the following instructions to install the dependencies.

### Cloning

```sh
$ git clone https://github.com/lucas625/Projeto-CG.git
$ cd project folder
```

### Golang

- Windows
  
Follow the download page on [golang website](https://golang.org/dl/).

- Linux

```sh
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```
- Checking Version

```sh
$ go version
```

### Dependencies

go mod will take care of the dependencies.

## Input

Once you run any script specified in the next section, you will have to choose an *.obj* file with triangularized faces and a few other files.

Files:

- The **Object** as an *.obj* file, please notice that there are a few examples available at *resorces/obj*.
- The **Camera** as an *.json* file, the camera.json is available at *resources/json* you just need to edit it, but notice that if if you set the vectors as an empty list then the camera will use *lookat* algorithm with the camera's position and the center point of the bounding box of the object as parameters. It would be wise to let the application find the camera by itself.
- The **Light** as an *.json* file, the light is available in the same folder that contains the **Camera**, but pay attention that you must specify all light's data for the scene.
- The **Object's Light Properties** as an *.json* file, the Object's Light Properties is available in the same folder that contains the **Camera**, but pay attention that you must specify all light's data for the object.

## Running the Project

- Testing

```sh
go run tests/test.go
```

- Running
  
Unavailable for now.