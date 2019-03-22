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
    - [Mateix](#mateix)
    - [MateixWatch](#mateix-watch)
- [Development](#development)
  - [Pre-Requisites](#pre-requisites)
  - [Setup](#setup)
  - [File Structure](#file-structure)
- [Resources](#resources)
- [License](#license)
- [Original Problem](#original-problem)

## Problem Statement

**We need to create a solution to synchronize two folders in two different devices, each folder have a file called `data`, which needs to be synchronized and it should be bi-directional**. To finally decide upon our solution, we need to make sure the following characteristics are met.

- **Reliability:** Solution should be reliable, and should take care of all the possible cases it can go wrong.

- **Automatic:** Everything should happen in the background with minimum human intervene.

- **Cheap:** It should do everything in minimum time while using very few resources, with least latency.

> Find the original problem statement at the end of the file.

## Challenges

Following are the challenges we need to tackle, mentioned with their possible solutions.

**Challenge: How should we communicate between two computers?**

When communicating over the internet, following possible scenarios:

No |IP-1 | IP-2 |  Details |
---|-----|------|----------|
 1 | Static | Static |`Server to Server` <br> Example: Backup servers connected to production servers.
 2 | Dynamic | Static | `Server to remote devices` <br> Example: Dropbox, Google drive connected to remote devices.
 3 | Dynamic | Dynamic | `Two remote devices` <br> Example: Two remote devices connected.

For our solution, we will only cover the 1st case. The same solution can be applied for other cases, which little changes. We have multiple methods to communicate over the internet, some are mentioned below.

No | Method | Secure | Reliable | Speed | Automatic |
---|--------|--------|----------|-------|-----------|
1 | TCP (Transmission Control Protocol) | Medium | High | Fast | Yes
2 | SSH (Secure Shell) | High | High | Medium | Not by default
3 | UDP (User Datagram Protocol) | Less | Less | Fastest | Yes

*Result*

After comparing all the methods, I have decided to use **`TCP`**.

***

**Challenge: When files should be synchronized?**

Following are the options on when files should be synchronized:

1. Immediately `inotify-tools`
2. After a time gap `crontab`

*Result*

After comparing all the options, I have decided to synchronize **`immediately`** a change is detected. *It will make it real time, will prevent merge conflict.*

> **Note:** if we update the device 1, updates will be sent to device 2, that will trigger the script, which will try to update the device 1, and this might go in a loop. We need to be handle the **butterfly effect**.

***

**Challenge: How should we measure the difference?**

To measure the difference, we can use any of method mentioned below.

No | Method | Reliable | Speed
---|--------|----------|------
1 | Time Modification (Metadata) | `Less Reliable` <br>software can and does manipulate the modification time. Also, the user might change system time and confuse the sync program. | `Fast` <br> O(1)
2 | Checksum (Hash the file) | `More Realible` <br> It's an (almost) certain way measure difference, hash collisions do happen, but It is rare. | `Slow` <br> O(n)

*Result*

After comparing all the methods, I have decided to use **`Checksum`**.

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

> **Note:** vice-versa is also true in this action table.

Following is the action table based on the time:

Time x | Time x+1 | Action
-------|----------|-------
Does not exist | Exist | Created
Existed | Does not exist |  Deleted
Exist | Modification | Modification

*Result*

For this solution, we are just tracking one file `data`, and we are doing synchronization immediately, so most of the cases won't apply to us.

***

**Challenge: How will you handle a merge conflict?**

Following are the methods we have to prevent merge conflict:

No | Method | Details | Merge Quality | Automatic
---|--------|---------|---------------|-----------
1 | Ask the user | Ask the user how to merge them or which one to pick. | Best| No
2 | Lock other user files | Lock a file if it is owned by the other user. | No Merge | Yes
3 | overwrite with latest changes | We can overwrite the file with latest changes. | Medium | Yes

> **Note:** Resolving merge conflict is technically imposible without human intervene

*Result*

Again, in our case, this won't be a constant problem, we have just one file. So I have decided to use `overwrite with latest changes` method.

***

## Final Solution

``` yml
STEP 1: Required tools are installed.
- inotify-tools

STEP 2: Binary and service files are downloaded, and service is enabled.
- /usr/bin/mateix
- /usr/bin/mateixWatch
- /etc/systemd/system/mateix-watch.service
- daemon-reload
- enable mateix-watch.service

STEP 3: Create the dotfiles and config files.
- /etc/.mateix/
- /etc/.mateix/syncList
- /etc/.mateix/log  

STEP 4: Creating a mateix watched folder.
- Add the path to /etc/.mateix/syncList
- Create PATH/.mateix/config.json

STRUCTURE OF `config.json`:
{
  targetIP: // Ip of other device
  targetDir: // Path of sync folder in other device
}

WORKING:
`inotifywait` is a tool part of `inotify-tools`, which is used
to catch changes in a folder or file.

`mateix-watch.service` is started when the system boot, which
activates the `mateixWatch` script.

`mateixWatch` to catch any changes using `inotifywait` in the files
and folders which are listed in `/etc/.mateix/syncList`.

Once changes are detected, `mateixWatch` send the $PATH to `mateix`
to do the synchronisation amd log it to `/etc/.mateix/log`.

`mateix` is the binary tool which reads the `$PATH/.mateix/config.json`
file to fetch `targetIP`, then it will connect to the server using the
`targetIP`, and fetch the checksum fo the `targetDir`, if the checksum
doesn't match, it will send the changes to the server, and server will
update the local `data` file. Checksum also help in solving the butterfly
effect.


```

## Usage

Using Mateix is very simple. All you have to do is install this CLI tool.

> Right now, the install script is tested only on a Debian-Based Distribution, but it can be easily configured for the other distros aswell.

#### Installation

To install Mateix, open your terminal, and type the commands given below.

1. Download the [Install](https://raw.githubusercontent.com/ramantehlan/mateix/master/install) script. `$ wget https://raw.githubusercontent.com/ramantehlan/mateix/master/install`
2. Make the script executables. `$ chmod +x ./install`
3. Execute the `install` script. `$ ./install`


#### Commands  

Once Mateix is installed, now you can use it to sync folders. Following are the commands available to help you sync folders.

##### Mateix

`Mateix` is the main program which will get the job done, it will communicate with the other systems, and synchronize files.

Command <br> (Prefix: `mateix`) | Working |
--------|---------|
init | To initialize a mateix watched folder.
update | To update a mateix watched folder. <br> `--file` to provide the path which needs to be updated.
server | `--start` to start the server. <br> `--stop` to stop the server. <br> Server use port `1248`
uninstall | To uninstall the mateix from the system

> **Note:** In any case, you must not rename your mateix watched folder. Since, it's location is added to /etc/.mateix/syncList, on rename it will misbehave.

##### mateixWatch

`mateixWatch` is the script which catch the changes in files, and call `mateix` command to take care of it. `mateixWatch` is automaticlly executed when the system starts by `mateix-watch.service`. It supports only following commands.

Command <br> (Prefix: `mateixWatch`) | Working
--------|---------|
start | To start the watch program on the files listed in `/etc/.mateix/syncList`. It use `inotifywait` to catch changes. <br><br> **Note:** I do not suggest using this command to start the watch program. Instead, you should start the `mateix-watch.service` service. If you still wish to use it, make sure you are a root user.
stop | To stop the watch program, by killing all the `inotifywait` processes.


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
├── init.go
├── install
├── LICENSE
├── main.go
├── mateix
├── packages
│   ├── command
│   │   └── command.go
│   └── e
│       └── e.go
├── README.md
├── server.go
├── service
│   ├── mateixWatch
│   └── mateix-watch.service
├── uninstall.go
└── update.go
```

No | File Name  | Purpose |
---|------------------|---------|
1 | `README.md` | Current file you are reading
2 | `LICENSE` | GNU GPL V3.0 License
3 | `install` | Install script
4 | `mateix` | Complied binary to run the tool
5 | `mateixWatch`  | Script to catch changes in folders
6 | `mateix-watch.service` | Service to call `mateixWatch` on boot
7 | `main.go` | Main/initial file of the program
8 | `init.go` | Initialize the mateixWatch in a repository
9 | `update.go` | Sync a folder with targetIP
10 | `uninstall.go` | Uninstall the mateix tool from system
11 | `server.go` | Server to get changes from client
12 | `command.go`| Package to help in execution of unix commands
13 | `e.go` | Package to handle errors

## Resources

- [Wiki File Synchronization](https://en.wikipedia.org/wiki/File_synchronization)
- [Go Lang Org](https://golang.org/)
- [TCP](https://en.wikipedia.org/wiki/Transmission_Control_Protocol)
- [SSH](https://en.wikipedia.org/wiki/Secure_Shell)
- [UDP](https://en.wikipedia.org/wiki/User_Datagram_Protocol)
- [pkg/net](https://golang.org/pkg/net/)
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
