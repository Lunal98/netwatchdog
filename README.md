# netwatchdog
monitor connection of headless linux devices and attempt to automatically remediate issues

# What it needs to accomplish

It needs to loops through a set of steps and attempt to remediate any given faults, as well as make debugging faults that may not be remediated by the software.
## monitoring loop
- check the interfaces
- check network access
- check the defined fstab mounts 
## check the interfaces
if interface exactly one interface should have an ip of 172.x.x.x or 192.x.x.x
if neither one has, attempt to reconnect to wifi
if both has, disconnect wifi
in both cases log the change between the previous checks interface config and the new one
## check network access
check network access by pinging the main server or the gateway
if the interface is still online but network is not available attempt to remediate by reconnecting network (after waiting a configured minutes)

## check the defined fstab mounts 
check each folder mount
	compare mount output with fstab?
	check if the folder actually exists/writable?
if missing try to remount

## modularity

define a way to add checks to the loops
provide an interface for these checks
each check needs to have a descriptive name, a function for the check, and a remediation process, as well as a fatality of the check (whether or not the os needs to be rebooted if the remediation steps fail to correct the state of the checked object).

## CLI

The Service needs to be interactable through a set of cli commands
## High level classes
**Check scheduler**


# Example of how it should be usable by a user

```Go
package main

import (
	"github.com/Lunal98/netwatchdogCore"
	"github.com/Lunal98/netwatchdogChecks"
	"examplecheck"
)

func main() {
	nwd := netwatchdogCore.New()
	Interface := InterfaceCheck{
		Interface: "eth0",
		Priority: 0,
	}
	Net := NetworkCheck{
		Subnet: "192.168.100.0/24",
		Ping: "8.8.8.8",
		Priority: 1,
	}
	SMB := SMBCheck{
		Share: "//dmhuftp/Invoices",
	  //Share: "/mnt/fileserver"
	  //Share: "auto"
		Priority: 2,
	}
	nwd.AddCheck(Interface,30*time.Second)
	nwd.AddCheck(Net,30*time.Second)
	nwd.AddCheck(SMB,30*time.Second)
	nwd.Start()

}
```