# Golang

- [Writing Web Applications - The Go Programming Language](https://golang.org/doc/articles/wiki/)

- [Build Web Applications With Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/preface.html)
   Chapters 2, 3, 7 might be useful.

- [Extremely useful documentation page](golang.org/pkg/net/http)

- [Pretty much the bible](https://golang.org/doc/effective_go.html)

# Documentation

Here are some resources for material each of you should know, as well as some timelines as to when you should know it.

# General Tooling

These are general tools that we'll be using throughout the project. You should definitely understand these before you start, or at least be able to figure out how to use them.

## git

Using git is absolutely imperative to our collaboration. If you're not already proficient with git, you should absolutely take some time to practice the basics. 

Check out [try.github.io](try.github.io) to practice the basics.

To go beyond the basics, you might want to check out the [the git book](https://git-scm.com/book/en/v2), which is the definitive source of information on git. Chapters 2 and 3 should be enough for the basics, and read chapter 1 if you want to have a deeper understanding of why git is important, what problems git solves, and how it differs from Subversion and earlier VCSs.

Please name branches descriptively. Name them to describe what changes the branch contains. __Do not__ name them "robbies-branch" or "my-branch2". It's really not helpful to other people, and in a week you'll completely have forgotten what the purpose of that branch was.

Secondly, do not have long-standing branches. You should make small changes on your branch before submitting a pull request. No branch should be in existence for more than a week (ideally). If you get stuck on something and your branch is living longer than a couple of days, please contact me for help. I'd love to help you overcome your problem and help you get your branch merged in.

## Vagrant

Vagrant is a tool we're using to virtualize our development environment. I use a Mac to code, and Justice and Connor use Windows. We don't want someone encountering a bug and another person saying "it works on my machine", because then which computer do we use as the source of truth? We want code that works on my Mac to work on Connor's PC. Consistency is important for obvious reasons.

Vagrant gives us consistency by allowing us to use an indentical virtual machine to develop our code in, such that it doesn't matter which OS you're using, because your code gets tested inside of the virutal machine, which is identical for all of us.

Vagrant is a command line tool to spin up and create virutal machines. It uses Virtualbox to run the virtual machines. Vagrant gives us a convenient way to install software onto the virtual machine, as well as package the machine for later use. 

Make sure you have Vagrant and Virtualbox installed.

## Minikube

This is another very important one. Minikube allows you 
