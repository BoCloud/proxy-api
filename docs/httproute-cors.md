##### Cors

- Cors用于客户端请求跨域配置
- 在HTTPRoute中Cors配置分别出现在两个位置 `Spec.Routes[0].Route.Cors 和 `Spec.Routes[0].Route.Rules[0].Cors`，如果在一个HTTRoute中配置都配置了Cors生效规则如下

- Ingress Controller将会采用就近原则，即优先`Spec.Routes[0].Route.Rules[0].Cors`，其次`Spec.Routes[0].Route.Cors`，若不设置则使用默认配置，Ingress Controller采用了更细致的处理方法，根据处理优先级会对Cors中的所有字段进行合并，然后再应用到Nginx配置

| 字段                 | 类型   | 必填 | 描述 | 示例 |
| -------------------- | ------ | ---- | ---- | ---- |
| CorsAllowOrigin      | string | 否   |      |      |
| CorsAllowMethods     | string | 否   |      |      |
| CorsAllowHeaders     | string | 否   |      |      |
| CorsAllowCredentials | string | 否   |      |      |
| CrosAllowCredentials | bool   | 否   |      |      |
| CorsExposeHeaders    | string | 否   |      |      |
| CorsMaxAge           | int    | 否   |      |      |



#### Cors默认值

