##### DefaultBackend

- HTTPRoute支持当后端服务返回错误码在`ErrorCode`中时，请求重定向到`DefaultService`

- DefaultBackend位于HTTPRoute的`Spec.Routes[0].Rules[0].DefaultBackend`

| 字段      | 类型                                                         | 必填 | 描述            | 示例       |
| --------- | ------------------------------------------------------------ | ---- | --------------- | ---------- |
| Service   | *[DefaultService](httproute-defaultbackend.md#defaultservice) | 否   | 指向Serivce配置 |            |
| ErrorCode | [int]                                                        | 否   | 错误码          | [503, 501] |

###### DefaultService

| 字段 | 类型   | 必填 | 描述        | 示例    |
| ---- | ------ | ---- | ----------- | ------- |
| Name | string | 否   | service名称 | default |
| Port | *int32 | 否   | service端口 | 80      |



###### 示例

- HTTPRoute

```yaml
apiVersion: proxy.bocloud.io/v1beta1
kind: HTTPRoute
metadata:
  name: httproute-defaultserver
spec:
  ingressClassName: nginx
  routes:
    - host: defaultserver.demo.com
      protocol: http
      rules:
        - path: /user
          pathType: exact
          backends:
            - name: service
              port: 80
          defaultBackend:
            service:
              name: default
              port: 8080
            errorCode: [501, 503]
```

- nginx.conf

```nginx
	## start server defaultserver.demo.com
	server {
		server_name defaultserver.demo.com ;
		
		listen 80  ;
		listen [::]:80  ;
		listen 443  ssl http2 ;
		listen [::]:443  ssl http2 ;
		
		location @custom_custom-default-backend-default-default_501 {
			internal;
			proxy_intercept_errors off;
			proxy_set_header       X-Code             501;
			proxy_set_header       X-Format           $http_accept;
			proxy_set_header       X-Original-URI     $request_uri;
			proxy_set_header       Host               $best_http_host;
			set $proxy_upstream_name "custom-default-backend-default-default";
			rewrite                (.*) / break;
			proxy_pass            http://upstream_balancer;
		}
		
		location @custom_custom-default-backend-default-default_503 {
			internal;
			proxy_intercept_errors off;
			proxy_set_header       X-Code             503;
			proxy_set_header       X-Format           $http_accept;
			proxy_set_header       X-Original-URI     $request_uri;
			proxy_set_header       Host               $best_http_host;
			set $proxy_upstream_name "custom-default-backend-default-default";
			rewrite                (.*) / break;
			proxy_pass            http://upstream_balancer;
		}
		
		location = /user {			
			error_page 501 = @custom_custom-default-backend-default-default_501;
			error_page 503 = @custom_custom-default-backend-default-default_503;
			proxy_pass http://upstream_balancer;	
		}
	}
```

------

​																					  [跳转HTTRoute](httproute.md)
