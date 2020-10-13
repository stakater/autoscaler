package deletenode

import (
    "context"

    apiv1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    kube_client "k8s.io/client-go/kubernetes"
)

func DeleteNode(node *apiv1.Node, client kube_client.Interface) error {
    return client.CoreV1().Nodes().Delete(context.TODO(), node.Name, metav1.DeleteOptions{})
}