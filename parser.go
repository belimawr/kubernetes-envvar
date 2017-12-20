package main

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

type deployment struct {
	Spec struct {
		Containers []struct {
			Name    string `yaml:"name"`
			EnvVars []struct {
				Key   string `yaml:"name"`
				Value string `yaml:"value"`
			} `yaml:"env"`
		} `yaml:"containers"`
	} `yaml:"spec"`
}

func parse(file []byte) ([]map[string]string, error) {
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
