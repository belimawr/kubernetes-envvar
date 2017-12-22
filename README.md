Kubernetes-envvar
=================
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

TODO
----

* Remove ugly workaround helm templates
* Write unit tests
* Parse deployment.yaml in a more generic way
* Review the logging/debug messages
* Return the return code of the command run

Licence
-------
GPLv3

