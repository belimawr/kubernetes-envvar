package main

import "testing"

func Test_parseDeployment(t *testing.T) {
	expectedLength := 1
	expectedMap := map[string]string{
		"FOO": "bar",
		"BAR": "foo",
	}

	m, err := parseDeployment([]byte(deploymentMock))

	if err != nil {
		t.Errorf("did not expect an error: %q", err)
	}

	if len(m) != expectedLength {
		t.Fatalf("expecting len(m) = %d, got %d", expectedLength, len(m))
	}

	for k, v := range expectedMap {
		if v != m[0][k] {
			t.Errorf("expecting m[%s] = %q, got: %q", k, v, m[0][k])
		}
	}
}

func Test_parseDeployment_parse_error(t *testing.T) {
	expectedLength := 0

	m, err := parseDeployment([]byte(invalidYAML))
	if err == nil {
		t.Errorf("expecting an error")
	}

	if len(m) != expectedLength {
		t.Fatalf("expecting len(m) = %d, got %d", expectedLength, len(m))
	}
}

func Test_parseDeploymentWithTemplate(t *testing.T) {
	expectedLength := 1
	expectedMap := map[string]string{
		"FOO":    "bar",
		"ANSWER": "42",
	}

	m, err := parseDeploymentWithTemplate([]byte(deploymentWithTemplateMock))

	if err != nil {
		t.Errorf("did not expect an error: %q", err)
	}

	if len(m) != expectedLength {
		t.Fatalf("expecting len(m) = %d, got %d", expectedLength, len(m))
	}

	for k, v := range expectedMap {
		if v != m[0][k] {
			t.Errorf("expecting m[%s] = %q, got: %q", k, v, m[0][k])
		}
	}
}

func Test_parseDeploymentWithTemplate_helm_workaround(t *testing.T) {
	expectedLength := 1
	expectedMap := map[string]string{
		"FOO":    "bar",
		"ANSWER": "42",
	}

	m, err := parseDeploymentWithTemplate(
		[]byte(deploymentWithTemplateHelmTemplateMock))

	if err != nil {
		t.Errorf("did not expect an error: %q", err)
	}

	if len(m) != expectedLength {
		t.Fatalf("expecting len(m) = %d, got %d", expectedLength, len(m))
	}

	for k, v := range expectedMap {
		if v != m[0][k] {
			t.Errorf("expecting m[%s] = %q, got: %q", k, v, m[0][k])
		}
	}
}

func Test_parseDeploymentWithTemplate_parse_error(t *testing.T) {
	expectedLength := 0

	m, err := parseDeploymentWithTemplate([]byte(invalidYAML))
	if err == nil {
		t.Errorf("expecting an error")
	}

	if len(m) != expectedLength {
		t.Fatalf("expecting len(m) = %d, got %d", expectedLength, len(m))
	}
}

var deploymentMock = `apiVersion: v1
kind: Pod
metadata:
  name: Test_parseDeployment
  labels:
    purpose: testing
spec:
  containers:
  - name:  Test_parseDeployment
    image: gcr.io/google-samples/node-hello:1.0
    env:
    - name: FOO
      value: "bar"
    - name: BAR
      value: "foo"
`

var deploymentWithTemplateMock = `apiVersion: v1
kind: Pod
metadata:
  name: Test_parseDeploymentWithTemplate
  labels:
    purpose: testing
spec:
  someKey: some value
  template:
    spec:
      containers:
        - name: Testing container
          image: gcr.io/google-samples/node-hello:1.0
          env:
            - name: FOO
              value: "bar"
            - name: ANSWER
              value: "42"
`

var deploymentWithTemplateHelmTemplateMock = `apiVersion: v1
kind: Pod
metadata:
  name: Test_parseDeploymentWithTemplate
  labels:
    purpose: testing
    broken_key: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
spec:
  someKey: some value
  template:
    spec:
      containers:
        - name: Testing container
          image: gcr.io/google-samples/node-hello:1.0
          env:
            - name: FOO
              value: "bar"
            - name: ANSWER
              value: "42"
`

var invalidYAML = `
""INVALID"`
