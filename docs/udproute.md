#### UDPRoute

- UDPRoute是Kubernetes中的一种自定义资源，在Kubernetes中使用UDPRoute来定义udp路由信息，Ingress Controller会监听该资源的创建、更新、删除等事件并同步更新Nginx配置。当在集群中创建UDPRoute，Ingress Controller会首先对该资源的配置进行检查，检查通过后则UDPRoute创建成功，Ingress Controller接收到UDPRoute的创建成功事件并更新Nginx配置。

- 文件位置： [udproute](../apis/proxy/v1beta1/udproute_types.go)

- 规定：udproute字段注释存在 +unsupported，表示Ingress Controller尚未支持该项配置


#### 资源定义

| 字段                               | 类型          | 必填 | 描述                                                         | 示例     |
| ---------------------------------- | ------------- | ---- | ------------------------------------------------------------ | -------- |
| Spec.IngressClassName              | *string       | 否   | IngressClass的名称，如果为空则使用默认IngressClass，用于指定哪个Controller处理此资源 | nginx    |
| Spec.Streams                       | array[Stream] | 是   | 多组udp路由配置                                              |          |
| Spec.Streams[0].Stream.Port        | int32         | 是   | 外部端口                                                     | 8001     |
| Spec.Streams[0].Stream.TLS         | *struct       | 否   | 用于需要tls认证的udp通信                                     | 尚未支持 |
| Spec.Streams[0].Stream.TLS.Secret  | string        | 否   | kubernetes中Secret名称                                       | 尚未支持 |
| Spec.Streams[0].Stream.ServiceName | string        | 是   | kubernetes中Service名称                                      | nginx    |
| Spec.Streams[0].Stream.ServicePort | string        | 是   | kubernetes中Service的port字段值                              | 8081     |
| Status.Conditions                  | array[struct] | 否   | 该CRD资源的基本状态信息，由控制器填充                        |          |


#### 示例
```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: UDPRoute
metadata:
  name: udproute-sample
spec:
  ingressClassName: nginx
  streams:
    - port: 8800
      serviceName: web
      servicePort: 80
```