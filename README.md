# Pucker

Pucker headlessly stands up a VoIP client upon the connection of a bluetooth device. It was originally built for Driver/Paddock communications for [The 24 Hours of Lemons.](https://24hoursoflemons.com/) When a bluetooth device is connected, components of system configurations trigger to sweep away all old bluetooth and audio connections and log into a [Mumble server.](https://wiki.mumble.info/wiki/Installing_Mumble) It exists because I began working on solving a problem I had with [talkiepi](https://github.com/dchote/talkiepi), which in turn came from [barnard](https://github.com/layeh/barnard).

## How's it work?

[Full documentation](https://docs.wotlemons.com/doku.php?id=r_d:comms)

As I have it configured pucker relies on an Ansible playbook colocated with the mumble server. It is silly to reach out to another machine to trigger the ansible run, but it also allows me as an operator to easily trigger the environment we want to stand up on the client. This way I can debug OTA if I see a driver having issues.

Pair your bluetooth devices. Devices should automatically connect to an available device, given they'll accept a connection and they're in range.

Install udev rules, and target script.

Install Ansible and place playbook.

Connect a device. Syslog output is very helpful and relevant output from each stage in the process will be logged.

Note: Though there are only 4 steps above, each of those steps has numerious intricacies you must master in order to pass through the Bridge of Hell unscathed.



## License

MPL 2.0

## Author

- pucker - [Tom McNulty](https://github.com/TLMcNulty)
- talkiepi - [Daniel Chote](https://github.com/dchote)
- Barnard,Gumble Author - Tim Cooper (<tim.cooper@layeh.com>)
