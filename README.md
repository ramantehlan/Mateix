# mateix
A simple file synchronisation tool.

> Mateix means 'same' in the Catalan language, which is a language of Catalonia community of Spain.

# Index

- [Problem Statement](#problem-statement)
- [Challenges](#challenges)
- [Final Solution](#final-solution)
- [Usage](#usage)
  - [Installation](#installation)
  - [Commands](#commands)
- [Development](#development)
  - [Pre-Requisites](#pre-requisites)
  - [Setup](#setup)
  - [File Structure](#file-structure)
  - [Working](#working)
- [Resources](#resources)
- [License](#license)
- [Original Problem](#original-problem)

## Problem Statement

The gist of it is. **We need to create a solution to synchronize two folders in two different devices**. To finally decide upon our solution, we need to make sure the following characteristics are met.

- **Reliability:** Solution should be reliable, and should take care of all the possible cases it can go wrong.

- **Secure:** The whole process should be secure.

- **Automatic:** Everything should happen in the background with minimum human intervene.

- **Cheap:** It should do everything in minimum time while using very few resources, with least latency.

> Find the secret problem statement at the end of the file.

## Challenges

Following are the challenges we need to tackle, mentioned with their possible solutions.

**Challenge: How should we communicate between two computers?**


Any selected solution should work effectively in the following possible scenarios:

IP-1 | IP-2 |  Details |
-----|------|----------|
Static | Static |`Server to Server` <br> Example: Backup servers connected to production servers.
Dynamic | Static | `Server to remote devices` <br> Example: Dropbox, Google drive connected to remote devices.
Dynamic | Dynamic | `Two remote devices` <br> Example: Two remote devices connected.

We have multiple methods to communicate securely over the internet, which will also work for the above scenarios. Following are the top 3 of selected methods:

1. SSH
2. Eternal Terminal
3. Mosh

After comparing all the methods, I have used **`Mosh`** since it fits the
needs best. It is *highly efficient, works well on low bandwidth/connection
is persist over different networks, works well on all the operating systems.*

***

**Challenge: When files should be synchronized?**

Following are the options on when files should be synchronized:

- Immediately

  To do this, we need a file watcher program, which will watch our files and will update us about any changes, this will not just take fewer resources, but is a better option as it will be real time. It will use more resources only when we are heavily using the system. This will also be better for the merge conflict as the files will be updated in real time.

  However this might cause one problem, if we update the device 2, a call will be triggered as well, which might send the changes back to device 1. This will probably cause a butterfly effect.

  Also, the problem with inotify tool is that, Inotify does not support recursively watching directories, so if the main folder have sub-folder, it won't track them.

  Rename events are not handled directly; i.e., inotify issues two separate events that must be examined and matched in a context of potential race conditions.

  `inotify-tools`

- After a time gap.

  To do this, we need to set up a cron job, which will run after a particular period of time, and then it will sync everything. It will be using more resources as it will run even when there are no changes.

  `crontab -e`

***

**Challenge: How should we measure the difference?**


To measure the difference, we can use any of method mentioned below.

1. Time modification (Metadata)

It is **less reliable** since software can and does manipulate the modification
time. Also, the user might change system time and confuse the sync program. But,
it is a **faster** way to check if the files have been updated.

2. Checksum (Hash the files)

It's an (almost) certain way measure difference, hash collisions do happen,
but It is rare and therefor **more reliable**. Though it is **slow**,
as the file size will grow, it will get slower.

After comparing all the methods, I have used **`Time modification`** as a measure to
look for difference, since *the possibility of something going wrong is very less,
and it is the fastest way to do so.*

***

**Challenge: How should we tackle the differences?**

Following are the 3 cases that we need to handle:

- The file exists on device 1, not on device 2
- The file exists on both devices and is identical
- The file exists on both devices and is different

Following is the action table:

File 1 | File 2 | Action
-------|--------|-------
Deleted | No Deleted | Delete
Deleted | Deleted | Nothing
No Change | No change | Nothing
Modification | No change | Use A
Modification | Modification | Merge

> Also note, vice-versa is also true in this action table.

Following is the action table based on the time:

Time x | Time x+1 | Action
-------|----------|-------
Does not exist | Exist | Created
Existed | Does not exist |  Deleted
Exist | Modification | Modification

***

**Challenge: How will you handle a merge conflict?**

If the file exists and is different, you have to ask the user how to merge them or which one to pick. Asking regular users how to merge files is a bad idea. (Asking developers how to merge files is usually a bad idea.)

Another way to prevent merge conflicts is to lock a file on machine A if it’s being written to on machine B. This prevents an application on machine A from modifying it at the same time.

 Merges almost always require manual intervention and will often be unresolvable (either the user won’t know what to do and will just overwrite one side, or the file format won’t support lines-of-text style merging).

Provide only the read access to the other device, only the original owner will have the write access

***

**Challenge: What if something still goes wrong with data?**

A program no matter how well written, will always have that 0.1% that it will fail. In a case like that, the most important thing is the data. I have created a automatic git commits system, which will look for changes every 10 minutes, and will automaticlly trigger a commit. You can also switch it off if you want.

***

## Final Solution

Here will go the details of the final solution

## Usage

Using Mateix is very simple. First, you need to install this tool, and then use it as a shell command.

> Right now, the install script will work only on a Debian-Based Distribution, but it can be easily configured for the other distros aswell.

#### Installation

To install Mateix, open your terminal, and type the commands given below.

1. Download the [https://github.com/ramantehlan/mateix/releases/download/v1/install](#) script. `$ wget https://raw.githubusercontent.com/ramantehlan/mateix/v1/install`
2. Make the script executables. `$ sudo chmod +x ./install`
3. Execute the `install` script. `$ ./install`

This will not just install Mateix in the bin file, but will also install all the dependencies like git, ssh, crontab etc. You can check out more details of it when the script is getting downloaded.

> In order to sync with the other system, you will need to install Mateix in it too.

#### Commands  

Once Mateix is installed, now you can use it to sync folders. Following are the commands available right now to help you sync folders.

> Also note, all commands have a fixed prefix `mateix`. Example: if a command is `init`, it must be executed as `mateix init`.

Command | Working |
--------|---------|
init | To set up a folder for sync
--help | Print all the commands

## Development

If you are interested in developing this project, feel free to read more about it below.

#### Pre-Requisites

If you are interested in the development, then here are some pre-requisites you need to have.

- Familiarity of Go language.
- Knowledge of Git
- Unix/Linux Terminal

#### Setup

To set up the development environment in your system:

1. Install Go in your system.
2. Fork this repo, and clone it in your workplace.

#### File Structure

```
.
├── LICENSE
└── README.md
```

No | File/Folder Name | Purpose |
---|------------------|---------|
1 | `README.md` | Current file you are reading
2 | `LICENSE` | GNU GPL V3.0 License
3 | `install` | Install script

#### Working

After you install Mateix, an executable is created in bin file, whose Dotfiles are store in the home folder.

## Resources

- [Wiki File Synchronization](https://en.wikipedia.org/wiki/File_synchronization)
- [Go Lang Org](https://golang.org/)
- [SSH - Secure Shell](https://en.wikipedia.org/wiki/Secure_Shell)
- [Eternal Terminal](https://mistertea.github.io/EternalTerminal/)
- [Mosh](https://mosh.org/)
- [Most Research Paper](https://mosh.org/mosh-paper.pdf)
- [Git - Version Control System](https://en.wikipedia.org/wiki/Git)
- [what could be better than ssh](https://medium.com/@grassfedcode/what-could-be-better-than-ssh-e69561ec1b83)
- [File Synchronization Algorithms](https://ianhowson.com/blog/file-synchronisation-algorithms/)

## License

GNU GPL V3.0

## Original Problem

```
Clone Wars 2.0
--------------

Prime Minister Lama Su,

I hope this letter finds you in the best of health.

The last batch of clones you built for us was faulty
and did not perform as expected (https://www.youtube.com/watch?v=b0DuUnhGBK4)

We unearthed some secrets about how the droid army was trained and hope that
you can use this information to make a better army this time around. With the
galaxy on the brink of another war, I cannot help but emphasize how much a
large discount will help the Republic in its efforts.

One of our allies came across these schematics in an abandoned base that shed some
light on the droid training exercises, master Yoda concluded that a pair of droids
undergo various kinds of battle simulations during which each droid records its
progress and learning in a force, currently unfamiliar to us, called "Data".
This force from both droids is then combined in a ritual called the
"Sync" resulting in both droids having an increased data force.

Please have a look at this schematic, your engineers may have better luck
decoding its mysteries.

            +----------------+                +----------------+
            |                |                |                |
            |   +--------+   |      Sync      |   +--------+   |
            |   |-|Data|-|   | +------------> |   |-|Data|-|   |
            |   +--------+   | <------------+ |   +--------+   |
            |                |                |                |
            |    Driod  A    |                |    Driod  B    |
            |                |                |                |
            +----------------+                +----------------+

May the force be with you.

- Sifo-Dyas


[....2 months later....]


Prime Minister Lama Su!,

I hope the army is coming along nicely. The force has given us more clarity in
the last few months. As it turns out, this "Data" that we were so worried about,
is just a method by which the droids store information about their experiences and
orders. Most importantly, the "Sync" ritual was just an exchange of files
from one droid to another in both directions. This is how their data force
increased after the ritual.

Master Windoo has been doing extensive research and has come up with a simplified
experiment to test if this training method can be implemented. He says that you
should start by figuring out how to synchronize data between a folder on one
device (say device A) and a folder on another device (say device B).
In addition to that, a change made to the data on one device should also be made
available to the other device as well. If we have a way to do this then we could
potentially improve the quality of the new clone army. I hope your engineers
are able to make sense of all of this information. Do write back to me if you
need more information.

Please share your method and implementation in great detail with us so
that it can be added to our records in the Jedi Temple. I wish you luck.

May the force be with you.

- Sifo-Dyas


                                 +---------------------+
                                 | What's going on here?|
                                 +------------------+--+
                                                    |
                                                    |
  _                                                 |
  \\                                                |
   \\_          _.-._                               |
    X:\        (_/ \_)     <------------------------+
    \::\       ( ==  )
     \::\       \== /
    /X:::\   .-./`-'\.--.
    \\/\::\ / /     (    l
     ~\ \::\ /      `.   L.
       \/:::|         `.'  `
       /:/\:|          (    `.
       \/`-'`.          >    )
              \       //  .-'
               |     /(  .'
               `-..-'_ \  \
               __||/_ \ `-'
              / _ \ #  |
             |  #  |#  |   B-SD3 Security Droid
          LS |  #  |#  |      - Front View -

(http://www.ascii-art.de/ascii/s/starwars.txt)
```
