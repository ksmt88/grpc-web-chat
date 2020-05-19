grpc-web-chat
====

This is a test app that chats in real time.

![demo](demo/chat_demo_en.gif)

## Dependencies
- grpc
- grpc-web
- envoy
- redis
- grpc-web-chat-front(https://github.com/ksmt88/grpc-web-chat-front)

## Getting Started
```bash
docker-compose up
```

## Kubernetes
- create ingress, service...etc
```
kubectl apply -f ./k8s/
```

## EKS
### Create repository
```bash
aws ecr create-repository --repository-name grpc-web-chat_backend
aws ecr create-repository --repository-name grpc-web-chat_proxy
```
#### Image push
##### login
```bash
export AWS_ACCOUNT_ID=$(aws sts get-caller-identity --output text --query 'Account' --profile [profile])
aws ecr get-login-password --profile [profile] | docker login --username AWS --password-stdin https://${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com
```
##### build
```bash
docker-compose build
```
##### backend
```bash
export BACKEND_IMG_VERSION=[version]
docker tag grpc-web-chat_backend:latest ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-web-chat_backend:${BACKEND_IMG_VERSION}  
docker push ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-web-chat_backend:${BACKEND_IMG_VERSION}  
```
##### proxy
```bash
export PROXY_IMG_VERSION=[version]
docker tag grpc-web-chat_proxy:latest ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-web-chat_proxy:${PROXY_IMG_VERSION}  
docker push ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-web-chat_proxy:${PROXY_IMG_VERSION}  
```

##### kubectl
```bash
kubectl expose deployment proxy --type=LoadBalancer --name=proxy-service -n chat
kubectl get service/proxy-service |  awk {'print $1" " $2 " " $4 " " $5'} | column -t
```

ref: https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html
