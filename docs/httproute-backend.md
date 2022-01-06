#### Backend

- 支持多个后端服务

##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule[0].Backend

| 字段                  | 类型              | 必填 | 描述                            | 示例     |
| --------------------- | ----------------- | ---- | ------------------------------- | -------- |
| Name                  | string            | 否   | Service名称                     |          |
| Port                  | *int32            | 否   | Service端口                     |          |
| Weight                | *int32            | 否   | 灰度权重                        |          |
| Matches               | array[HTTPMatch]  | 否   | 匹配规则                        |          |
| Strategy              | string            | 否   | 会话保持                        |          |
| ChangeCookieOnFailure | bool              | 否   |                                 |          |
| FailTimeOutSeconds    | *int              | 否   |                                 | 尚未支持 |
| MaxFails              | *int              | 否   |                                 | 尚未支持 |
| MaxConns              | *int              | 否   |                                 | 尚未支持 |
| Keepalive             | *int              | 否   |                                 | 尚未支持 |
| Options               | map[string]string | 否   | [options](httproute-options.md) | 尚未支持 |



##### Spec.Routes[0].Route.Rules[0].HTTPRouteRule[0].Backend[0].HTTPMatch

- 匹配规则

| 字段     | 类型   | 必填 | 描述                    | 示例   |
| -------- | ------ | ---- | ----------------------- | ------ |
| Type     | string | 是   | 可选值: header / cookie | header |
| GroupId  | int    | 是   | 分组                    | 1      |
| Key      | string | 是   | key名称                 |        |
| Operator | stirng | 是   | 操作                    |        |
| Value    | strng  | 是   | 值                      |        |

