/*

kubernetes-envvar is a cli application to export the environment variables from
Kubernetes' deployment.yamal.

kubernetes-envvar reads and exports the environment variables in Kubernetes'
deployment.yaml and them and runs a command in the "new" environment

Usage:

	kubernetes-envvar <type: d|t> <filepath> <command and args>

Arguments:
 type: d|t
   d: means a simple deployment file (see example/deployment.yaml)
   t: means a helm deployment.yaml using template (see example/deploymentWithTemplate.yaml)

 filepath:
   is the deployment.yaml filepath

 command and args:
   is the command to be run and its args. They're going to be passed directly to exec.Command


*/
package main
