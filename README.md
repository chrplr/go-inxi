go-inxi:  get information about your computer system (os, hardware)
===================================================================

This is a port to Go of the Perl tool [inxi](https://github.com/smxi/inxi) which reports information about you computer system.

For example, here is the ouput of ./inxi-go ran on my laptop:

```
Machine:    product: Precision 5490  System: Dell Inc.  Type: laptop
System:     Host: is158520  Kernel: 6.17.0-19-generic x86_64  Uptime: 7h 48m
            OS: Ubuntu 25.10  Shell: bash 5.2.37  Desktop: ubuntu:GNOME
CPU:        Model: Intel(R) Core(TM) Ultra 7 165H  Info: 16 cores / 22 threads  Speed: 1003 MHz (min: 400 / max: 4700)
Memory:     RAM: total: 30.04 GiB  used: 9.61 GiB (32.0%)  Swap: total: 8.00 GiB  used: 0 KiB (0.0%)
Graphics:   Card: Intel Corporation Meteor Lake-P [Intel Arc Graphics]  Driver: i915
            Card: NVIDIA Corporation AD107GLM [RTX 2000 Ada Generation Laptop GPU]  Driver: nvidia
Audio:      Card: sof-soundwire  Driver: snd_soc_sof_sdw
            Server: PipeWire  v: 1.4.7  ALSA: k6.17.0-19-generic
```

The port was essentially performed by Claude Sonnet 4.6. 

Not all functionnalities of the original inxi were implemented (by far!) because we tried to remain as cross platform as possible.

Binaries for different platforms are available from the [releases](https://github.com/chrplr/go-inxi/releases).

If macOS pretends no to be able to run them check out [this page](https://chrplr.github.io/note-about-macos-unsigned-apps/)


Christophe Pallier, April 2026

---

License: GNU General Public License v3 or later — see [LICENSE.txt](LICENSE.txt)

(Respecting the original inxi perl script LICENSE). 

---


I could not resist copying here an from inxi's README, which I approve:

### APPLE CORPORATION OSX

Non-free/libre OSX is in my view a BSD in name only. It is the least Unix-like 
operating system I've ever seen that claims to be a Unix, its tools are mutated, 
its data randomly and non-standardly organized, and it totally fails to respect 
the 'spirit' of Unix, even though it might pass some random tests that certify a 
system as a 'Unix'. 

If you want me to use my time on OSX features or issues, you have to pay me, 
because Apple is all about money, not freedom (that's what the 'free' in 'free 
software' is referring to, not cost), and I'm not donating my finite time in 
support of non-free operating systems, particularly not one with a market 
capitalization hovering around 1 trillion dollars, with usually well north of 
100 billion dollars in liquid assetts. 


### MICROSOFT CORPORATION WINDOWS


To be quite clear, support for Windows will never happen, I don't care about 
Windows, and don't want to waste a second of my time on it. I also don't care 
about cygwin issues, beyond maybe hyper basic issues that can be handled with a 
line or two of code. inxi isn't going to ruin itself by trying to handle the 
silly Microsoft path separator \, and obviously there's zero chance of my trying 
to support PowerShell or whatever else they come up with. 

While I would consider doing Apple stuff if you paid my hourly full market 
rates, in advance, I would not consider touching Windows for any amount of 
money. My best advice there is, fork inxi, and do it yourself if you want it. 
You'll soon run screaming from the project however, once you realize what a 
nightmare you've stepped into.

If you are interested in something like inxi for Windows, I suggest, rather than 
forking inxi, you just start out from scratch, and build the features up one by 
one, that will lead to much better code.

