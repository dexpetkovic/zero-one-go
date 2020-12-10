# Zero, One, GO!
-----
# Getting Started - An introduction into setting up your development environment

First lesson will cover basics of GoLang and guide you into setting up your development environment.

## Go Language Specification

Go is a general-purpose language designed with systems programming in mind. (from [GoLang spec](https://golang.org/ref/spec))

It is:
* strongly typed, where types (String, Integer, etc.) need to be defined or (automatically) inferred for attributes and methods
* garbage-collected, so that developers do not need to take care (too) much about memory management
* has explicit support for concurrent programming

Go programs are created from packages, which are building blocks of Go applications. We will deep dive into packages within this tutorial.

## Configuring Development Environment

As this is an opinionated guide, we will use VisualStudio Code within this tutorial.

The complete guide to get started is [in this Youtube Video]((https://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE))

A quick summary of steps you need to get started:
* Install Git from [project's website](https://git-scm.com/downloads)
* Install Microsoft VisualStudio Code from [project's website](https://code.visualstudio.com/download)
* From VisualStudio Code Extensions menu, install Go Extension `golang.go`
* Clone this project to your home directory with `git clone git@github.com:dexpetkovic/zero-one-go.git`

## What are we actually building?

The goal of this tutorial is to write a client library in Go which will access a REST API for management of bank accounts.

We will first implement our library against mocked API and later on use the real API.

Further to this:
* We will not implement authentication
* Bank accounts will be created, fetched, listed and deleted
* We will use these [API specifications](api-specs.md)
