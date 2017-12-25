Kubernetes-envvar
=================
[![Build Status](https://travis-ci.org/belimawr/kubernetes-envvar.svg?branch=master)](https://travis-ci.org/belimawr/kubernetes-envvar)


Small cli application to read environment variables from
Kubernetes' deployment.yaml, export them and run a command
in the "new" environment

**This is a WORK IN PROGRESS!!!**

Motivation
----------
I've been using Kubernetes for a while now and using environment variables
to configure my applications and each environment (staging, production, etc)
has got its own deployment.yaml and when the number of environment variables
grows past two it gets annoying to run an application locally pointing to a
different environment (e.g. staging).

So I decided to write this wee tool to do the "monkey job" of reading
the deployment.yaml and exporting the environment variables

Usage
-----

```sh
$ ./kubernetes-envvar <type: d|t> <filepath> <command and args>
```
Arguments:
* Type: It's either `t` or `d` depending on the structure of the yaml file.
  * `d` is for the simpler yaml (see [example/deployment.yaml](example/deployment.yaml))
  * `t` is for a helm yaml file (see [example/deploymentWithTemplate.yaml](example/deploymentWithTemplate.yaml))
* Filepath: Is the path to the Kubernetes' deployment.yaml
* Command: Is the command and it's args.

TODO
----

* Remove ugly workaround helm templates
* Write unit tests
* Parse/replace environment variables on the command and it's arguments
* Parse deployment.yaml in a more generic way
* Review the logging/debug messages
* Return the return code of the command run

Licence
-------
GPLv3

