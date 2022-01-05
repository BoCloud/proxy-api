#### 概念：

- tcproute是Kubernetes中的一种自定义资源，在Kubernetes中使用tcproute来定义tcp路由信息，Nginx Controller会监听该资源的创建、更新、删除等事件并同步更新Nginx配置。当在集群中创建tcproute，Nginx Controller会首先对该资源的配置进行检查，检查通过后则tcproute创建成功，Nginx Controller接收到tcproute的创建成功事件并更新Nginx配置。

- 定义文件位置： [tcproute](../apis/proxy/v1beta1/tcproute_types.go)

- 规定：定义文件中字段注释存在 +unsupported，表示Nginx Controller尚未支持该项配置

#### 定义

| 字段                | 类型         | 必填 | 解释                                                         | 示例     |
| ----------------------- | ------------ | ---- | ------------------------------------------------------------ | -------- |
| Spec.IngressClassName   | *string| 否   | IngressClass的名称，如果为空则使用默认IngressClass，用于指定哪个Controller处理此资源 | nginx    |
| Spec.Streams            | array[Stream] | 是   | 多组tcp路由配置                                              |          |
| Spec.Streams[0].Stream.Port        | int32        | 是   | 外部端口                                                     | 8001     |
| Spec.Streams[0].Stream.TLS         | *struct | 否   | 用于需要tls认证的tcp通信                                     | 尚未支持 |
| Spec.Streams[0].Stream.TLS.Secret  | string       | 否   | kubernetes中Secret名称                                       | 尚未支持 |
| Spec.Streams[0].Stream.ServiceName | string       | 是   | kubernetes中Service名称                                      | nginx    |
| Spec.Streams[0].Stream.ServicePort | string       | 是   | kubernetes中Service的port字段值                              | 8081     |
| Status.Conditions       | array[struct] | 否   | 该CRD资源的基本状态信息，由控制器填充                        |          |

#### 示例YAML

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: TCPRoute
metadata:
  name: tcproute-sample
spec:
  ingressClassName: nginx
  streams:
    - port: 8443
      serviceName: nginx
      servicePort: 443
```

