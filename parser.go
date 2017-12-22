package main

import (
	"log"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

type deployment struct {
	Spec struct {
		Containers []container `yaml:"containers"`
	} `yaml:"spec"`
}

type deploymentWithTemplate struct {
	Spec struct {
		Replicas string `yaml:"replicas"`
		Template struct {
			Spec struct {
				Containers []container `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

type container struct {
	Name    string    `yaml:"name"`
	EnvVars []envVars `yaml:"env"`
}

type envVars struct {
	Key   string `yaml:"name"`
	Value string `yaml:"value"`
}

func parseDeployment(file []byte) ([]map[string]string, error) {
	data := deployment{}

	if err := yaml.Unmarshal(file, &data); err != nil {
		log.Printf("could not unmarshal: %q", err.Error())
		return []map[string]string{}, err
	}

	variables := []map[string]string{}
	for _, container := range data.Spec.Containers {
		vars := map[string]string{}

		for _, envVar := range container.EnvVars {
			vars[envVar.Key] = envVar.Value
		}
		variables = append(variables, vars)
	}

	return variables, nil
}

func parseDeploymentWithTemplate(file []byte) ([]map[string]string, error) {
	data := deploymentWithTemplate{}

	if err := yaml.Unmarshal(file, &data); err != nil {
		// Ungly workaround helm templates
		log.Printf("got an error parsing helm file, trying to clean it...")
		re := regexp.MustCompile(`{{.*}}`)
		cleanFile := re.ReplaceAllString(string(file), "")

		if err := yaml.Unmarshal([]byte(cleanFile), &data); err != nil {
			log.Printf("could not unmarshal: %q", err.Error())
			return []map[string]string{}, err
		}
		log.Print("done!")
	}

	variables := []map[string]string{}
	for _, container := range data.Spec.Template.Spec.Containers {
		vars := map[string]string{}

		for _, envVar := range container.EnvVars {
			vars[envVar.Key] = envVar.Value
		}
		variables = append(variables, vars)
	}

	return variables, nil
}
