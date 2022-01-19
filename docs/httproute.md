#### HTTPRoute

- HTTPRoute是Kubernetes中的一种自定义资源，在Kubernetes中使用HTTPRoute来定义http路由信息，Ingress Controller会监听该资源的创建、更新、删除等事件并同步更新Nginx配置。当在集群中创建HTTPRoute，Ingress Controller会首先对该资源的配置进行检查，检查通过后则HTTPRoute创建成功，Ingress Controller接收到HTTPRoute的创建成功事件并更新Nginx配置。

- Nginx对七层协议的功能支持非常强大，也因此HTTPRoute的设计相对TCPRoute/UDPRoute复杂的多，HTTPRoute的字段结构设计在充分考虑了Nginx功能特性情况下扩展了多种灰度发布方案的支持。

- 文件位置： [httproute](../apis/proxy/v1beta1/httproute_types.go)

- 规定：文件中字段注释存在 +unsupported，表示Ingress Controller尚未支持该项配置

  

------

​    `Nginx对于http协议的支持非常丰富，HTTPRoute在充分考虑了Nginx配置的基础上，扩展了多种灰度方案支持。也因此HTTPRoute的设计结构相对复杂的多，HTTPRote支持配置多域名、每个域名下多路径、每个路径可以配置不同的灰度发布服务。在路径、服务等又可配置多种可选参数，因此关于HTTPRoute这块的字段解析，会拆分成多个文档分别讲述`

------



##### Spec

| 字段             | 类型                          | 必填 | 描述                                                         | 示例  |
| ---------------- | ----------------------------- | ---- | ------------------------------------------------------------ | ----- |
| IngressClassName | *string                       | 否   | IngressClass的名称，如果为空则使用默认IngressClass，用于指定哪个Controller处理此资源 | nginx |
| Routes           | [[Route](httproute.md#route)] | 是   | 多组http路由配置                                             |       |

##### Route  

- Spec.Routes[0]

| 字段     | 类型                                               | 必填 | 描述                                                         | 示例           |
| -------- | -------------------------------------------------- | ---- | ------------------------------------------------------------ | -------------- |
| Host     | string                                             | 否   | 域名FQDN                                                     | demo.nginx.com |
| Protocol | string                                             | 否   | Nginx转发给后端使用的协议，可选值HTTP、HTTPS、GRPC，默认值HTTP；尚未支持GRPC | HTTP           |
| TLS      | *[TLS](httproute.md#tls)                           | 否   | Host的tls证书，用于Nginx验证请求来源是否合法                 |                |
| Proxy    | *[Proxy](httproute-proxy.md)                       | 否   | 详细信息见Proxy文档                                          |                |
| Cors     | *[Cors](httproute-cors.md)                         | 否   | 详细信息见Cors文档                                           |                |
| Rules    | [[HTTPRouteRule](httproute.md#httprouterule)]      | 是   | 多组http路由                                                 |                |
| Options  | [Options](httproute-options.md#specroutes0options) | 否   | 详细信息见Options文档                                        |                |

##### TLS

- Spec.Routes[0].TLS

| 字段   | 类型   | 必填 | 描述                                                    | 示例        |
| ------ | ------ | ---- | ------------------------------------------------------- | ----------- |
| Secret | string | 是   | 存储tls证书的secret名称，需要和此CRD在同一个namespace下 | secret-name |

##### HTTPRouteRule

- Spec.Routes[0].Rules[0]

| 字段           | 类型                                                     | 必填 | 描述                                         | 示例     |
| -------------- | -------------------------------------------------------- | ---- | -------------------------------------------- | -------- |
| Path           | string                                                   | 是   | 详情见[Path](httproute-path.md)              | /user    |
| PathType       | string                                                   | 否   | 路径匹配方式exact、prefix、regex，默认prefix | exact    |
| Rewrite        | string                                                   | 否   | 路径重写，支持普通以及正则等                 | /abc     |
| Proxy          | *[Proxy](httproute-proxy.md)                             | 否   | 详细信息见Proxy文档                          |          |
| Cors           | *[Cors](httproute-cors.md)                               | 否   | 详细信息见Cors文档                           |          |
| RateLimit      | [Ratelimit](httproute-ratelimit.md)                      | 否   | 详细信息见Ratelimit文档                      |          |
| Options        | [Options](httproute-options.md#specroutes0rules0options) | 否   | 详细信息见Options文档                        | 尚未支持 |
| Backends       | [[Backend](httproute-backend.md)]                        | 否   | 详细信息见Backend文档                        |          |
| DefaultBackend | *[DefaultBackend](httproute-defaultbackend.md)           | 否   | 详细信息见DefaultBackend文档                 |          |

#### Status

| 字段     | 类型   | 必填 | 描述                                  | 示例 |
| -------- | ------ | ---- | ------------------------------------- | ---- |
| Hostname | string | 否   | 该CRD资源的基本状态信息，由控制器填充 |      |



#### 示例

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-sample
spec:
  ingressClassName: nginx
  routes:
    - host: ingress.demo.com
      tls:
        secret: tls-secret
      protocol: http
      rules:
        - path: /user
          rateLimit:
            connections: 3
            rpm: 3
            rps: 5
          pathType: exact
          cors:
            corsMaxAge: 10000
          backends:
            - name: service
              port: 80
        - path: /app
          rateLimit:
            connections: 3
            rpm: 3
            rps: 5
          pathType: exact
          backends:
            - name: service2
              port: 8080

---
apiVersion: v1
kind: Secret
metadata:
  name: tls-secret
  namespace: default
data:
  tls.crt: string
  tls.key: string
type: kubernetes.io/tls

```