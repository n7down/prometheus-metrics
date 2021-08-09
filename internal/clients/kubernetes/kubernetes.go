package kubernetes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/n7down/prometheus-metrics/internal/clients/kubernetes/responses"
)

type KubernetesClient struct {
	apiServer string
	token     string
}

func NewKubernetesClient(apiServer, token string) *KubernetesClient {
	return &KubernetesClient{
		apiServer: apiServer,
		token:     token,
	}
}

func (c *KubernetesClient) GetNumberOfPods(namespace string) int {
	var (
		response = responses.GetPodsResponse{}
		bearer   = fmt.Sprintf("Bearer %s", c.token)
	)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/namespaces/%s/pods", c.apiServer, namespace), nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return len(response.Items)
}
