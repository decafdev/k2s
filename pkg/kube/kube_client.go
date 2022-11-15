package kube

import (
	"context"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	netV1 "k8s.io/api/networking/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// // ClientInterface - interface
// type ClientInterface interface {
// 	ApplyDeploymentSpec(namespace string, deployment *appsV1.Deployment) (*appsV1.Deployment, error)
// }

// Client struct
type Client struct {
	k8s kubernetes.Interface
	// ext       apiext.Interface
	// dynamic   dynamic.Interface
	// discovery discovery.DiscoveryInterface
	ctx context.Context
}

// NewKubeService function description
func NewKubeClient() (*Client, error) {
	client := &Client{
		ctx: context.Background(),
	}
	return client.Connect()
}

// Connect method
func (c *Client) Connect() (*Client, error) {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	configPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return c, errors.Wrap(err, "failed to create K8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return c, errors.Wrap(err, "failed to create K8s clientset")
	}

	// apiextClient, err := apiext.NewForConfig(config)
	// if err != nil {
	// 	return t, errors.Wrap(err, "failed to create api extension client")
	// }

	// dynamicClient, err := dynamic.NewForConfig(config)
	// if err != nil {
	// 	return t, errors.Wrap(err, "failed to create dynamic client")
	// }

	// discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	// if err != nil {
	// 	return t, errors.Wrap(err, "failed to create discovery client")
	// }

	c = &Client{
		k8s: clientset,
		ctx: context.Background(),
		// ext: apiextClient,
		// dynamic:   dynamicClient,
		// discovery: discoveryClient,
	}

	return c, nil
}

// // ApplyCRD method
// func (t *Client) ApplyCRD(crd *apiExtV1Beta1.CustomResourceDefinition) (*apiExtV1Beta1.CustomResourceDefinition, error) {
// 	create := t.ext.ApiextensionsV1beta1().CustomResourceDefinitions().Create
// 	update := t.ext.ApiextensionsV1beta1().CustomResourceDefinitions().Update

// 	if res, err := create(t.ctx, crd, metaV1.CreateOptions{}); apierrors.IsAlreadyExists(err) {
// 		return update(t.ctx, crd, metaV1.UpdateOptions{})
// 	} else {
// 		return res, err
// 	}
// }


// ApplyServiceAccount method
func (c *Client) ApplyServiceAccount(namespace string, spec *coreV1.ServiceAccount) (*coreV1.ServiceAccount, error) {
	create := c.k8s.CoreV1().ServiceAccounts(namespace).Create
	update := c.k8s.CoreV1().ServiceAccounts(namespace).Update

	res, err := create(c.ctx, spec, metaV1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}

	return res, err
}

// ApplyNamespace method
func (c *Client) ApplyNamespace(namespace *coreV1.Namespace) (*coreV1.Namespace, error) {
	return c.k8s.CoreV1().Namespaces().Create(c.ctx, namespace, metaV1.CreateOptions{})
}

// ApplyClusterRole method
func (c *Client) ApplyClusterRole(spec *rbacV1.ClusterRole) (*rbacV1.ClusterRole, error) {
	create := c.k8s.RbacV1().ClusterRoles().Create
	update := c.k8s.RbacV1().ClusterRoles().Update

	res, err := create(c.ctx, spec, metaV1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}

	return res, err
}

// ApplyClusterRoleBinding method
func (c *Client) ApplyClusterRoleBinding(spec *rbacV1.ClusterRoleBinding) (*rbacV1.ClusterRoleBinding, error) {
	create := c.k8s.RbacV1().ClusterRoleBindings().Create
	update := c.k8s.RbacV1().ClusterRoleBindings().Update

	res, err := create(c.ctx, spec, metaV1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}

	return res, err
}

// ApplySecret method
func (c *Client) ApplySecret(namespace string, spec *coreV1.Secret) (*coreV1.Secret, error) {
	create := c.k8s.CoreV1().Secrets(namespace).Create
	update := c.k8s.CoreV1().Secrets(namespace).Update

	res, err := create(c.ctx, spec, metaV1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}

	return res, err

}

// ApplyDeployment method
func (c *Client) ApplyDeployment(namespace string, spec *appsV1.Deployment) (*appsV1.Deployment, error) {
	create := c.k8s.AppsV1().Deployments(namespace).Create
	update := c.k8s.AppsV1().Deployments(namespace).Update

	res, err := create(c.ctx, spec, metaV1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}

	return res, err
}

// ApplyIngress method
func (c *Client) ApplyIngress(namespace string, spec *netV1.Ingress) (*netV1.Ingress, error) {
	create := c.k8s.NetworkingV1().Ingresses(namespace).Create
	update := c.k8s.NetworkingV1().Ingresses(namespace).Update
	res, err := create(c.ctx, spec, metaV1.CreateOptions{})

	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}
	return res, err
}

// ApplyConfigMap method
func (c *Client) ApplyConfigMap(namespace string, spec *coreV1.ConfigMap) (*coreV1.ConfigMap, error) {
	create := c.k8s.CoreV1().ConfigMaps(namespace).Create
	update := c.k8s.CoreV1().ConfigMaps(namespace).Update

	res, err := create(c.ctx, spec, metaV1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return update(c.ctx, spec, metaV1.UpdateOptions{})
	}

	return res, err
}

// ApplyService method
func (c *Client) ApplyService(namespace string, spec *coreV1.Service) (*coreV1.Service, error) {
	get := c.k8s.CoreV1().Services(namespace).Get
	create := c.k8s.CoreV1().Services(namespace).Create
	update := c.k8s.CoreV1().Services(namespace).Update

	res, err := get(c.ctx, spec.ObjectMeta.Name, metaV1.GetOptions{})
	if apierrors.IsNotFound(err) {
		return create(c.ctx, spec, metaV1.CreateOptions{})
	}

	spec.ObjectMeta.ResourceVersion = res.ObjectMeta.ResourceVersion
	spec.Spec.ClusterIP = res.Spec.ClusterIP
	return update(c.ctx, spec, metaV1.UpdateOptions{})
}

// // GetSecret method
// func (t *Service) GetSecret(name, namespace string) (*coreV1.Secret, error) {
// 	return t.k8s.CoreV1().Secrets(namespace).Get(t.ctx, name, metaV1.GetOptions{})
// }

// // ListSecrets method
// func (t *Service) ListSecrets(namespace string) (*coreV1.SecretList, error) {
// 	return t.k8s.CoreV1().Secrets(namespace).List(t.ctx, metaV1.ListOptions{})
// }

// // GetConfigMap method
// func (t *Service) GetConfigMap(name, namespace string) (*coreV1.ConfigMap, error) {
// 	return t.k8s.CoreV1().ConfigMaps(namespace).Get(t.ctx, name, metaV1.GetOptions{})
// }

// // OnNamespaceEvent method
// func (t *Service) OnNamespaceEvent(options metaV1.ListOptions) rxgo.Observable {
// 	return rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
// 		watcher, err := t.k8s.CoreV1().Namespaces().Watch(t.ctx, options)
// 		if err != nil {
// 			ch <- rxgo.Error(err)
// 		}

// 		for event := range watcher.ResultChan() {
// 			ch <- rxgo.Of(event)
// 		}
// 	}})
// }

// // OnDeploymentEvent method
// func (t *Service) OnDeploymentEvent(name, namespace, release string) rxgo.Observable {
// 	timeout := int64(90)
// 	return rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
// 		watcher, err := t.k8s.AppsV1().Deployments(namespace).Watch(t.ctx, metaV1.ListOptions{
// 			TypeMeta: metaV1.TypeMeta{
// 				Kind: "Deployment",
// 			},
// 			LabelSelector:        fmt.Sprintf("release=%s", release),
// 			FieldSelector:        "",
// 			Watch:                true,
// 			AllowWatchBookmarks:  false,
// 			ResourceVersion:      "",
// 			ResourceVersionMatch: "",
// 			TimeoutSeconds:       &timeout,
// 			Limit:                0,
// 			Continue:             "",
// 		})
// 		if err != nil {
// 			ch <- rxgo.Error(err)
// 		}

// 		for event := range watcher.ResultChan() {
// 			ch <- rxgo.Of(event)
// 		}

// 		close(ch)
// 	}})
// }

// // CopySecret method
// func (t *Service) CopyRegistry(name, fromNamespace, toNamespace string) (registry *coreV1.Secret, err error) {
// 	source, err := t.GetSecret(name, fromNamespace)
// 	if err != nil {
// 		return registry, err
// 	}

// 	destinationCopy, err := t.ApplySecret(toNamespace, &coreV1.Secret{
// 		ObjectMeta: metaV1.ObjectMeta{
// 			Name:      source.ObjectMeta.Name,
// 			Namespace: toNamespace,
// 			Labels:    source.ObjectMeta.Labels,
// 		},
// 		Data: source.Data,
// 	})

// 	if err != nil {
// 		return registry, err
// 	}

// 	return destinationCopy, nil
// }

// // CreateRegistrySecret method
// func (t *Service) CreateRegistrySecret(o *CreateRegistryDTO) (*coreV1.Secret, error) {
// 	registry, err := NewContainerRegistry(&ContainerRegistryOptions{
// 		Name:      o.Name,
// 		Namespace: o.Namespace,
// 		Registry:  o.Registry,
// 		Username:  o.Username,
// 		Password:  o.Password,
// 	})

// 	if err != nil {
// 		return &coreV1.Secret{}, err
// 	}

// 	registry.Secret.ObjectMeta.Labels = map[string]string{
// 		"secret-type": "private-registry",
// 	}

// 	return t.ApplySecret(o.Namespace, registry.Secret)
// }

