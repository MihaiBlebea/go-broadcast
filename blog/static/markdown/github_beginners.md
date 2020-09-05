---
Title: Github for absolute beginners - How to get started with your first repo?
Summary: If you just started to code, then git is not for you... That's what they told me. But does it still apply today?
Image: /static/img/github.jpeg
Tags:
    - github
    - repo
    - coding
    - devops
Slug: github-for-beginners-first-repository
Published: "2020-08-27 10:04:05"
---
"If you just started to code, then git is not for you..."

This statemenent is so wrong.

You should start using git as soon as possible in your coding career.

**In fact:**

> The sooner you start using git, the faster you will progress over time.

It's not just about pushing your code to a place where it can be shared.

## This is how git helps you be a better developer

I remember in the "good old days" when I started coding, how I used to be so afraid that I will mess up my latest application, that I would create a copy of the folder, zip it and save it in another location on my laptop.

Soon enought I ended up with `application_v1.zip`, `application_v2.zip`, `application_latest.zip` and how can I forget the `application_latest_latest.zip` 😂😂😂.

Safe to say that eventually I lost my mind trying to figure out which latest is actually the real latest version 😂.

Then I discovered Github...

<img src="https://media0.giphy.com/media/13xHqoOQOdFu5a/giphy.gif" />

This are just a couple of points on how git can help you:

- Share your code with your friends, work, interviewers, managers, etc (even with your cat 😂 - if the cat is into coding)

- Version your code and eliminate the fear of messing things up - you are free to break the code as you can always go back to a previous version

- Become more disciplined - it's easy to code, but to be a great developer you have to learn how to follow the rules of your organization. All companies use some sort of git tool (Github, Bitbucket, Gitlab, etc)

If I managed to convince you to start using git, the next question should be:

> Where do I start?

I created a simple step by step guide that will get you from just absolute beginner, to atually creating your first remote repository and pushing code into it with just a couple of simple commands in the terminal.

But first... 

## How does git works?

A git repository is a place where you can safely hold your code and version it.

A repository can have many branches. 

Think of branches as different highways that lead to the same place, but all in a different way.

For now, we will just use the default `master` branch.

Every branch is usually composed of 2 parts:

- the remote part - hosted on a different server (Github, Gitlab, Bitbucket, etc)

- the local part - which is hosted on your laptop

We will start by adding code to our local master branch on our own computer, then push this local branch to the remote location to keep it in a safe place.

## Step 1 - Install git tool on your laptop

Before we start with the fancy stuff, let's make sure that we have all the dependencies installed on our machine.

For this tutorial we won't need much, just `git` a command line tool to interact with github (or any other git provider).

Check if you have it already installed with this command:

- `git --version`

You should get something simillar to this ```git version 2.23.0``` if git is already installed on your machine.

If not, let's install it using `homebrew` or `https://git-scm.com`.

### Homebrew

Homebrew is a dependency manager for Mac OS. Think **composer for PHP** or **npm for Node**.

If you don't have `homebrew` installed, I suggest you start using it, because it will make your life much easier.

You can find more information for installing homebrew here: 

<a href="https://brew.sh/" target="_blank">Install homebrew</a>

Or just run this in the terminal: ```/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"```

Now that you have homebrew installed, all you have to do is type this command to install git:

```brew install git``` - it's that simple.

### Using the git-scm.com website

If for some reason you don't feel confident enought to install the git cli using homebrew, then just navigate to the official website `https://git-scm.com`.

Under the downloads section you will see the download links for different operating systems. Just click on yours and follow the instructions.

Don't forget at the end to verify the installation with ```git --version``` command.

## Step 2 - Create your Github account (skip this step if you already have one)

Go to <a href="https://github.com" target="_blank">Github.com</a> and create an account. Just follow the instructions and don't forget to confirm your email address.

Once you have your new shinny account, go to the **+** icon in the top right corner and select `New repository` - keep in mind that this might change in the future, but for now when I write this article it's still there 😂.

You will be directed to this page where you can select a name for your repository and a bunch of other options.

Let's keep it simple for now, so just add a name and click the confirmation button.

<img src="/static/img/github_account_step_1.png" />

Next, you will see a bunch of instructions on how to push your code to github.

<img src="/static/img/github_account_step_2.png" />

You can see in the top left corner - the username (your profile name) and the name of the repository. 

That will be a unique slug that we will use in the future to get the code from github.

For now, just copy the line that starts with `git remote add origin ...`.

We will use this to instruct our git command line application to push the code to this repository. 

You may notice your github repository slug in that command.

## Step 3 - Push your awesome code to Github

Finally we reached the fun part.

Let's start pushing our code to Github.

First, use the terminal to navigate to the root folder of your application.

Then run this command to initialize a local git repository:

- ```git init```

You may notice a new folder in your project `.git`. This will be used by github to keep track of the changes in your local code and sync with the remote repository.

There are 3 steps to pushing code to github:

- The staging - select the files that you want to add to your repository

- The commit - create a bundle of your files and assign a description to the change that you want to make

- The push - push your change to github when you are happy with the changes made

Let's break this down into actionable steps.

### The staging 

Staging is very important because it allows you to fine tune which files you want to add to your repository.

Let's say you modify the files `foo.go`, `bar.go` and `config`.

But config was modified just to allow for a new configuration to run on your local laptop and you don't want that change to be part of your final repository.

So we will add just the first 2 files to staging:

- ```git add foo.go```

- ```git add bar.go```

If for some reason you want to add all your files to stage - this is usually usefull when you first create a new repository, just run:

- ```git add .```

The dot means add all the files.

One more important thing to consider here. If you want to exclude one or more files permanently from your staging, then create a `.gitignore` file in the root folder of your project (notice that there is a dot in front of the file name).

Add all the files or folders that you don't want to keep track of in that file and add the file to the staging.

When you are happy with your stage, you can make a last check with this command:

- ```git status```

You will get informations about what files are staged and which are not.

If you accidentaly added a file to stage that you shouldn't have, just run this command to exclude it:

- ```git restore foo.go``` - this will exclude the `foo.go` file from the stage.

### The commit

The next step is to commit the staged files.

This will create a bundle of all our changes and prepare them to be pushed to github.

Keep in mind that not all the commits needs to be pushed right away.

You can choose to create 3 or 5 commits before deciding that is time to push them to the main repository.

All the commits need to have a name - a description that tells a story of what are the changes made by this commit.

Usually a good rule of thumb is to use a tone that speaks about the change. For example:

- `Adds change to the interface`

- `Changes how the service works`

Notice that you can add `The commit...` in front and this will read `The commit adds change to the interface`.

This will make it much easier to keep track of the changes.

Just avoid commits like `Friday change` or `Changed the thing that I talked with Paul` because this will not make any sense over time.

To create your first commit, just use this command after you added the files to the stage:

- ```git commit -m "First commit"```

Notice how I used `First commit` here - this is how you should describe your first ever commit to a new repository. It's not a rule, just a choice.

### Linking your local repository with the github one that we created

This step is necesary only the first time when you create a new repository and you want to push your code.

Before you push the code, we will need to link our local repo with the github one that we just create earlier.

If you remember I asked you to copy that line that starts with `git remote add origin ...`.

Just paste it in your terminal and press enter.

Job done.

### Push your commit to the remote repository

Now that we have the local and the remote repository linked, we will carry on with pushing our code.

We won't get much into branching strategies (we will keep that for another article).

For now just use this command to push your changes:

- ```git push origin master```

This will tell git to push the code to the origin (the remote repository in Github) and to the `master` branch.

You will get a confirmation message at the end.

Congrats, you pushed your first code to github.

<img src="/static/img/yeey.gif" />

## Step 4 - Bonus - bundle everything into a single command

I don't suggest you do this, mostly because every step is important and you should always check what you commit to your remote repository.

But just to prove that it's possible to do this with one single command, we will add all the above steps into a `Makefile`.

First, you will need to have a github repository and link it with your local one
Create a file named `Makefile` in the root of your project

Add the code below to the file

```Makefile
git: add commit

add:
	git add .

commit:
	git commit -m "$(msg)"

push:
	git push origin master
```

After you create the Makefile and add the above content, just run this command in your root folder:

- ```make git msg="First commit"```

That's it!

In part 2 we will talk more about Github actions and how to create a true CI / CD pipeline that will build, test and deploy your code. All of this FREE by just using Github actions.
