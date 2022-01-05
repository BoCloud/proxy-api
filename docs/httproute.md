#### HTTPRoute

- HTTPRoute是Kubernetes中的一种自定义资源，在Kubernetes中使用HTTPRoute来定义http路由信息，Ingress Controller会监听该资源的创建、更新、删除等事件并同步更新Nginx配置。当在集群中创建HTTPRoute，Ingress Controller会首先对该资源的配置进行检查，检查通过后则HTTPRoute创建成功，Ingress Controller接收到HTTPRoute的创建成功事件并更新Nginx配置。

- Nginx对七层协议的功能支持非常强大，也因此HTTPRoute的设计相对TCPRoute/UDPRoute复杂的多，并且在充分考虑了Nginx功能特性情况下扩展支持多种灰度发布方案。

- 文件位置： [httproute](../apis/proxy/v1beta1/httproute_types.go)

- 规定：文件中字段注释存在 +unsupported，表示Ingress Controller尚未支持该项配置



#### 资源定义

##### 主体字段：

- httproute支持配置多个域名、域名下多个path、每个path又有不同的配置且每个path支持不同的灰度服务，所以其配置非常复杂
- 我们拆成几部分来讲述各个字段的功能含义，如是静态配置会具体到该配置会应用到Nginx的哪个配置字段

| 字段                            | 类型                   | 必填 | 描述                                                         | 示例           |
| ------------------------------- | ---------------------- | ---- | ------------------------------------------------------------ | -------------- |
| Spec.IngressClassName           | *string                | 否   | IngressClass的名称，如果为空则使用默认IngressClass，用于指定哪个Controller处理此资源 | nginx          |
| Spec.Routes                     | array[Route]           | 是   | 多组http路由配置                                             |                |
| Spec.Routes[0].Route.Host       | string                 | 否   | 域名FQDN                                                     | demo.nginx.com |
| Spec.Routes[0].Route.Protocol   | string                 | 否   | Nginx转发给后端使用的协议，可选值HTTP、HTTPS、GRPC，默认值HTTP；尚未支持GRPC | HTTP           |
| Spec.Routes[0].Route.TLS        | *struct[TLS]           | 否   | Host的tls证书，用于Nginx验证请求来源是否合法                 |                |
| Spec.Routes[0].Route.TLS.Secret | string                 | 是   | 存储tls证书的secret名称，需要和此CRD在同一个namespace下      | tls-cert       |
| Spec.Routes[0].Route.Proxy      | *struct[Proxy]         | 否   | Nginx转发代理配置                                            |                |
| Spec.Routes[0].Route.Cors       | *struct[Cors]          | 否   | Nginx跨域配置                                                |                |
| Spec.Routes[0].Route.Rules      | *struct[HTTPRouteRule] | 是   | 自定义http路由                                               |                |
| Spec.Routes[0].Route.Options    | map[string]string      | 否   | 自定义的可选项配置                                           |                |
| Status.Hostname                 | string                 | 否   | 该CRD资源的基本状态信息，由控制器填充                        | 尚未支持       |



#####   Spec.Routes[0].Route.Proxy

- Nginx作为Http请求代理转发服务，其代理服务本身有很多可配置项，该列表描述了所有可配置项以及其字段含义解释
- 在httproute中分别有两个位置存在Proxy字段，生成Nginx配置阶段会采用就近原则，优先使用rules中配置，其次Route中Proxy，再次使用默认配置

| 字段                 | 类型   | 必填 | 功能描述 | 示例 |
| -------------------- | ------ | ---- | -------- | ---- |
| BodySize             | string | 否   |          |      |
| ConnectTimeout       | int    | 否   |          |      |
| SendTimeout          | int    | 否   |          |      |
| ReadTimeout          | int    | 否   |          |      |
| BuffersNumber        | int    | 否   |          |      |
| BufferSize           | string | 否   |          |      |
| CookieDomain         | string | 否   |          |      |
| CookiePath           | string | 否   |          |      |
| NextUpstream         | string | 否   |          |      |
| NextUpstreamTimeout  | int    | 否   |          |      |
| NextUpstreamTries    | int    | 否   |          |      |
| ProxyRedirectFrom    | string | 否   |          |      |
| ProxyRedirectTo      | string | 否   |          |      |
| RequestBufering      | string | 否   |          |      |
| ProxyBuffering       | string | 否   |          |      |
| ProxyHTTPVersion     | string | 否   |          |      |
| ProxyMaxTempFileSize | string | 否   |          |      |



##### Spec.Routes[0].Route.Cors

- 支持用户自定义请求Cors
- 在httproute中分别有两个位置存在Cors字段，生成Nginx配置阶段会采用就近原则，优先使用rules中配置，其次Route中Cors，再次使用默认配置

| 字段                 | 类型   | 必填 | 功能描述 | 示例 |
| -------------------- | ------ | ---- | -------- | ---- |
| CorsAllowOrigin      | string | 否   |          |      |
| CorsAllowMethods     | string | 否   |          |      |
| CorsAllowHeaders     | string | 否   |          |      |
| CorsAllowCredentials | string | 否   |          |      |
| CrosAllowCredentials | bool   | 否   |          |      |
| CorsExposeHeaders    | string | 否   |          |      |
| CorsMaxAge           | int    | 否   |          |      |



##### Spec.Routes[0].Route.Options

- 可扩展的自定义字段

| 字段             | 类型 | 必填 | 功能描述 | 示例 |
| ---------------- | ---- | ---- | -------- | ---- |
| upstream-streams | bool | 否   |          |      |
|                  |      |      |          |      |


##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule

- 定义http路由信息，应用于Nginx location阶段

| 字段                  | 类型                    | 必填 | 功能描述                                     | 示例     |
| --------------------- | ----------------------- | ---- | -------------------------------------------- | -------- |
| Path                  | string                  | 是   | 路径                                         | /user    |
| PathType              | string                  | 否   | 路径匹配方式exact、prefix、regex，默认prefix | exact    |
| Rewrite               | string                  | 否   | 路径重写，将接收到的path重写转发到后端       | /abc     |
| Proxy                 | *struct[Proxy]          | 否   | 转发代理配置                                 |          |
| Cors                  | *struct[Cors]           | 否   | 跨域配置                                     |          |
| RateLimit             | struct[ReteLimit]       | 否   | 限速配置                                     |          |
| RateLimit.Connections | int                     | 否   |                                              |          |
| RateLimit.RPM         | int                     | 否   |                                              |          |
| RateLimit.RPS         | int                     | 否   |                                              |          |
| Options               | map[string]string       | 否   |                                              | 尚未支持 |
| Backends              | array[Backend]          | 否   | 支持配置多个后端服务进行灰度转发             |          |
| DefaultBackend        | *struct[DefaultBackend] | 否   | 默认服务，截获后端                           |          |


##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule[0].Backend

| 字段                  | 类型              | 必填 | 功能描述    | 示例     |
| --------------------- | ----------------- | ---- | ----------- | -------- |
| Name                  | string            | 否   | Service名称 |          |
| Port                  | *int32            | 否   | Service端口 |          |
| Weight                | *int32            | 否   | 灰度权重    |          |
| Matches               | array[HTTPMatch]  | 否   | 匹配规则    |          |
| Strategy              | string            | 否   | 会话保持    |          |
| ChangeCookieOnFailure | bool              | 否   |             |          |
| FailTimeOutSeconds    | *int              | 否   |             | 尚未支持 |
| MaxFails              | *int              | 否   |             | 尚未支持 |
| MaxConns              | *int              | 否   |             | 尚未支持 |
| Keepalive             | *int              | 否   |             | 尚未支持 |
| Options               | map[string]string | 否   |             | 尚未支持 |


##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule[0].DefaultBackend


| 字段         | 类型            | 必填 | 功能描述        | 示例           |
| ------------ | --------------- | ---- | --------------- | -------------- |
| Service      | *DefaultService | 否   | 指向Serivce配置 |                |
| Service.Name | string          | 否   | Service名称     | defaultService |
| Service.Port | *int32          | 否   | Service端口     | 80             |
| ErrorCode    | array[int]      | 否   | 错误码          | [503, 501]     |



##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule[0].Backend[0].HTTPMatch


| 字段     | 类型   | 必填 | 功能描述                | 示例   |
| -------- | ------ | ---- | ----------------------- | ------ |
| Type     | string | 是   | 可选值: header / cookie | header |
| GroupId  | int    | 是   | 分组                    | 1      |
| Key      | string | 是   | key名称                 |        |
| Operator | stirng | 是   | 操作                    |        |
| Value    | strng  | 是   | 值                      |        |


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
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURMVENDQWhXZ0F3SUJBZ0lVYks1bFZTYWhRLzRZTm5HWkJPekI5dEhlZXpFd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0pqRVJNQThHQTFVRUF3d0libWRwYm5oemRtTXhFVEFQQmdOVkJBb01DRzVuYVc1NGMzWmpNQjRYRFRJeApNVEV5TXpBME5EUXdPVm9YRFRJeU1URXlNekEwTkRRd09Wb3dKakVSTUE4R0ExVUVBd3dJYm1kcGJuaHpkbU14CkVUQVBCZ05WQkFvTUNHNW5hVzU0YzNaak1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0MKQVFFQXluMUluU3NBNHhoNFhyT28zU2lHRWp5S1BxeG5KeWtvU2hDbXdady9ZUUc2bmJUbXFPazBEU2FVOSs3TQpOWUVoZTE1Uk01ZGliNzVEd1A4TTBOMmQ3QlUyMHZCU215WVp4YU9sVG1NbmRNMFBzaDBnUzg4NG5DNlNnelA1CjVJS2F3TTgxdUMvU09PY3ZOdW53L1l5OVZRVzlIVEFiRkhOSzFuNDhJTE9vT20xVUZVSUJLRXpTV2VyZ3NNYjgKbzQzenpVRUxWSmJEVmRueDFtSVRsdzdVTksyNU1XeXZZaVRKMHJwaWRMUVhibitHcmdMRndkZW9FL3RVVlpCVwo0STByQVNxUDZuR3EwR28zWnJNRlRzZjhQVGZ4eDhNeTB6STQvQ2ZaSEtvTjl5ajhzdzJsbFlDTlp3MmdHT0V1CjlRUkk1ZXhtb3RUYmNaN0J6ZHZLaEx0aGxRSURBUUFCbzFNd1VUQWRCZ05WSFE0RUZnUVU4NWZacEs3Smx5L1AKSElzb1EvNnI1VEw3MzY4d0h3WURWUjBqQkJnd0ZvQVU4NWZacEs3Smx5L1BISXNvUS82cjVUTDczNjh3RHdZRApWUjBUQVFIL0JBVXdBd0VCL3pBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQWdhNzNSZk9xSVg2QkJWbjVyYk5jCkVrNjlRUXVzK2o0N2R1b2ZROUlWaWQ5MG91ZjZhc0NPa241ZnNXZjdzNEZYSEErUWhLR2xYaU81VWxFTkQreUQKK1ZzaGZKcFQzeFR3UmJ0WjRPWVhHeGJXcE01MGJUenJYWm9qa2c0M0hIczYrL1VIQXNWTW9uemZ4bmZTd1lISAp3N1lhbE16OUVJenpCQytxUzRVelAyZnc1SjYybkZ4WkxWUUt6YlVPcDNNcEpnL0xDTEk2TS9weTJtRTlpbTZVCkhPZDV2bjltcEhmYXFsTnNvVEhHazNBQ3dmRnBUK1R1bGpYb3laR3o1TGI3eHBNcVl0YVgrOHZZMitpOEhkbzEKYUM4eFY1OC9GQ3BVUjV6bzY2TzdwVncyRzFmdkFQY1RQYnkzalAwcmtYd2RzZUErNThpN2h5Q3BBOHprUWxSRQpHdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
  tls.key: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRREtmVWlkS3dEakdIaGUKczZqZEtJWVNQSW8rckdjbktTaEtFS2JCbkQ5aEFicWR0T2FvNlRRTkpwVDM3c3cxZ1NGN1hsRXpsMkp2dmtQQQovd3pRM1ozc0ZUYlM4RktiSmhuRm82Vk9ZeWQwelEreUhTQkx6emljTHBLRE0vbmtncHJBenpXNEw5STQ1eTgyCjZmRDlqTDFWQmIwZE1Cc1VjMHJXZmp3Z3M2ZzZiVlFWUWdFb1ROSlo2dUN3eHZ5ampmUE5RUXRVbHNOVjJmSFcKWWhPWER0UTByYmt4Yks5aUpNblN1bUowdEJkdWY0YXVBc1hCMTZnVCsxUlZrRmJnalNzQktvL3FjYXJRYWpkbQpzd1ZPeC93OU4vSEh3ekxUTWpqOEo5a2NxZzMzS1B5ekRhV1ZnSTFuRGFBWTRTNzFCRWpsN0dhaTFOdHhuc0hOCjI4cUV1MkdWQWdNQkFBRUNnZ0VBQmxDNS93emtUakRwTUNyeVRWT0NPdmRnYUd3QUc1eVJBUjViMVJZR2RBUVYKeWUxbWRFWXh0V2RLcGlEd2hZcXRmS2VJYU0rRDVuQk10S3cvdmhQclpQMlVaQ2ZTcTd3WWVhMk03bER4WGhjMwpNaHJ1Y3U3WG1TZHFzbVRnbWx2b2I3TUd2ZVBmN3A1blBwTTFUUE1peEpBVlFkL0tPRzBRSEhoN2I1bXEyWWVaCklXL3BBMHYvdUt1T056RWNFTnY3YyttTFBlRFdpNW5zV20xZFlMRXFjemVnUG5IRDZrWXBPcTg3bzFPZnZVc2EKM3p6NnVvYlVVTVgvOWVKM2FFZkNJVmE1alRHS08zaUlLVC82bW8vRUlGM3daK2ExTTQ0MDJuYlZCbGllRWFrVApnUC9pd2g0OERWcC9yT2FiVWZBUXpJT05CU2NLVHRJTUVUbGJGZkR6MFFLQmdRRDVBTHg0UEI3ZXRSalVlcmJECmFEVTlka3c0R1ZLeFkrSld1akpzWFk4NnBld3ZIY0M2K2ZpUkN2MlhXTzJxS0V4cEtzendMaTR0V3RqTFpFTlEKRE5lZDY5OWFSRnI1Q05CaG9FL3I2enJUK3B6VElEVWQ2UHlUSHpZUXlQdFg2OW55K1pIZzVpMWY1VWdjOVhvSApjQmw5b29NZStLZnJ3NS9kbnN1SytzM2pCd0tCZ1FEUUxmRHdPVFFXSUtjYVpUMW5rNE5Cbi9sd25ISWJhS1B5CnZNajZJU0lHT05Ra25yRXJnbmJ1QjAzSlNQU2lRMWFwL25nMzFKZ1hGQVhqSWhmQzJkTk8wdzY2NjFaa0FGM0QKaEQyTFN6VW1QVlpqang0NXRDMlV2dXpXTk9yR3lzb2lBelFvb1N4VW9vUkJsZGRZakFIL1N1SFMyVUtMd0h6Rwp0NmNCQVNmamd3S0JnUURSZlNZVFBmbDJ3d08xMTl3bGdHbXlZUEYxRFJEK1B4dXdmWXhva1Rvc1RHWHRxZWw1CkpVOVRyOXgwVlpQMllWc1A2N3Rwb21DbE5kWkpIL3hsdjdnem03dFl0VU9ZV3lyOXg2TVZ4OXpCZFFvMXNkWWUKYU9MK1gvYmJua3VmeDhTZzRBazBIbE0wWjdFSTlCbUxZbXQvd0pieUdwOGtBbnhnTnZYbDRtWVBSd0tCZ1FDWAo1RDA4ZVBCSkNOQURrVVNKTXZiOHhjVVE1Z1RYZkxUS1lmWGRrcGtwb3dNZUtPOHB5TW9QaUNLNEwxUFdwSDB5CkVTb1R5amloOWdrSm5SRnJLTlZsV05jUmlLNEN6c1dhNXZ5a2lsNGdKWGJIczEraFNKWk5SalMxWWV5KzJLMDgKdmN1cnJWVVQ5M082Q3FNUnh5Mlo0RC8rUUdpdVlPWnBjd3dWem9zVkV3S0JnRWV3ZUJoSUQ1U251b2c1VjZYYQpWS0FRVXNHTld6MFA3RUt4VUZLOU1Wbm9lcFZzT1BnMVJTZFhGUmVIbitCRGtvU0JieW9nY1JvbEx3Q05kVUFWCm9NUm13QU03MWhMelRaVDBlT2tqZmFVUDBneFM5TjFpb2VQSmh2emlYTVBCbFBtWlNtWlNxUUM4ZnlISmJCNlIKMktyeWtLTXNONEgzbTh6Qkx2MmxVL3FUCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
kind: Secret
metadata:
  name: tls-secret
  namespace: default
type: kubernetes.io/tls
```