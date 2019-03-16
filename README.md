# mateix
An easy file synchronisation tool.

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
- [Resources](#resources)
- [License](#license)

## Problem Statement

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

The gist of it is. **We need to create a solution to synchronize two folders in two different devices**. To finally decide upon our solution, we need to make sure the following characteristics are met.

- **Reliability:** Solution should be reliable, and should take care of all the possible cases it can go wrong.

- **Secure:** The whole process should be secure.

- **Automatic:** Everything should happen in the background with minimum human intervene.

- **Cheap:** It should do everything in minimum time while using very few resources, with least latency.

## Challenges

Following are the challenges we need to tackle, mentioned with their possible solutions.

**Challenge 1:**

How should we communicate between two computers?

**Solution:**

Possible case

IP-1 | IP-2 | Example | Details |
-----|------|---------|---------|
Static | Static | Server to Server | *Details*
Dynamic | Static | Server to remote devices | *Details*
Dynamic | Dynamic | Two remote devices | *Details*

How to handle the dynamic address?
using UUID to identify the machine, since the IP keeps changing

**Challenge 2:**

How should we measure the difference?

**Solution:**

We can use any of the below methods to check if the files are changed.
- time modification
- Checksum
- Or both of them?

Possible cases

1. The file exists on device 1, not on device 2
2. The file exists on both devices and is identical
3. The file exists on both devices and is different

File 1 | File 2 | Action
-------|--------|-------
Deleted | No Deleted | Delete
Deleted | Deleted | Nothing
No Change | No change | Nothing
Modification | No change | Use A
Modification | Modification | Merge

Time x | Time x+1 | Action
-------|----------|-------
Does not exist | Exist | Created
Existed | Does not exist |  Deleted
Exist | Modification | Modification

**Challenge 3:**

How will you handle a merge conflict between the same file

**Solution:**

**Challenge 4:**

How and after what time span should the files be sync?

**Solution:**

Below are the possible two possible cases?
- Immediately?
- Update over a period of time?
  - Crone job? `crontab -e`

**Challenge 5:**

What if something still goes wrong with data?

**Solution**

Sync only one folder

- Use of git

**Challenge 7:**
How to set up everything?

**Solution:**

Dotfiles in the home folder, executables in the bin file. Also possibly each folder can have a dot file too, to store the metadata.


## Final Solution

Here will go the details of the final solution

## Usage

Using Mateix is very simple. First, you need to install this tool, and then use it as a shell command.

#### Installation

To install Mateix, open your terminal, and type the commands given below.

1. Download the [install](#) script. `$ wget https://github.com/something/something/install`
2. Make the script executables. `$ sudo chmod +x ./install`
3. Execute the `install` script. `$ ./install`

This will not just install Mateix in the bin file, but will also install all the dependencies like git, ssh, crontab etc. You can check out more details of it when the script is getting downloaded.

> In order to sync with the other system, you will need to install Mateix in it too.

#### Commands  

Once Mateix is installed, now you can use it to sync folders. Following are the commands available right now to help you sync folders.

> Also note, all commands have a fixed prefix `mateix`. Example: if a command is `init`, it must be executed as `mateix init`.

Command | Working |
--------|---------|
$ mateix init | To setup a folder for sync

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


## Resources

- [Wiki File Synchronization](https://en.wikipedia.org/wiki/File_synchronization)
- [Go Lang Org](https://golang.org/)
- [SSH - Secure Shell](https://en.wikipedia.org/wiki/Secure_Shell)
- [Git - Version Control System](https://en.wikipedia.org/wiki/Git)

## License

GNU GPL V3.0
