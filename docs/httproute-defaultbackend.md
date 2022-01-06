#### DefaultBackend

##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule[0].DefaultBackend

- 默认后端服务

| 字段         | 类型            | 必填 | 描述            | 示例           |
| ------------ | --------------- | ---- | --------------- | -------------- |
| Service      | *DefaultService | 否   | 指向Serivce配置 |                |
| Service.Name | string          | 否   | Service名称     | defaultService |
| Service.Port | *int32          | 否   | Service端口     | 80             |
| ErrorCode    | array[int]      | 否   | 错误码          | [503, 501]     |

