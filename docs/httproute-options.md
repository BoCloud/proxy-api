##### Options

- 为了提高HTTPRoute的可扩展性，因此增加了Options字段，目前尚未完善
- Options出现在HTTPRoute的`Spec.Routes[0].Options`、`Spec.Routes[0].Rules[0].Options`和`Spec.Routes[0].Rules[0].Backends[0].Options`
- 这三个位置的Options之间并无继承关系，每个都是独立的处理逻辑

###### Spec.Routes[0].Options

| 字段   | 类型   | 必填   | 描述   | 默认值 |
| ------ | ------ | ------ | ------ | ------ |
| 不支持 | 不支持 | 不支持 | 不支持 | 不支持 |



###### Spec.Routes[0].Rules[0].Options

| 字段   | 类型   | 必填   | 描述   | 默认值 |
| ------ | ------ | ------ | ------ | ------ |
| 不支持 | 不支持 | 不支持 | 不支持 | 不支持 |



###### Spec.Routes[0].Rules[0].Backends[0].Options

| 字段             | 类型   | 必填 | 描述                                                         | 默认值 |
| ---------------- | ------ | ---- | ------------------------------------------------------------ | ------ |
| service-upstream | string | 否   | 默认情况下会Nginx会轮训Pod IP转发请求<br>当配置为"true"时使用ClusterIP代替Pod IP/端口 | 无     |



##### 示例

- HTTPRoute

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-options
spec:
  ingressClassName: nginx
  routes:
    - host: options.demo.com
      protocol: http
      options:
        support: "no"
      rules:
        - path: /user
          pathType: exact
          options:
            support: "no"
          backends:
            - name: service
              port: 80
              options:
                service-upstream: "true"
        - path: /app
          pathType: exact
          backends:
            - name: service2
              port: 8080
```
