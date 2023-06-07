package kubernetes

// This go file contains the Dockerfile templates for different languages

var deployment = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: <Image>
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: <Port>`

var service = `apiVersion: v1
kind: Service
metadata:
  name: myapp
spec:
  selector:
    app: myapp
  ports:
  - port: <Port>
    targetPort: <Target Port>`

var ingress = `apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: myingress
  labels:
    name: myingress
spec:
  rules:
  - host: <Host>
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: <Service>
            port: 
              number: <Port>`

var secret = `apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  password: <Password> # This should be base64 encoded
`

var configmap = `apiVersion: v1
kind: ConfigMap
metadata:
  name: myapp
data:
  key: value`

var pod = `apiVersion: v1
kind: Pod
metadata:
  name: myapp
  labels:
    name: myapp
spec:
  containers:
  - name: myapp
    image: <Image>
    resources:
      limits:
        memory: "128Mi"
        cpu: "500m"
    ports:
      - containerPort: <Port>`

var pv = `apiVersion: v1
kind: PersistentVolume
metadata:
  name: mypv
spec:
  capacity:
    storage: <Size>
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  storageClassName: slow
  mountOptions:
    - hard
    - nfsvers=4.1
  nfs:
    path: /tmp
    server: 172.17.0.2`

var pvc = `apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mypvc
spec:
  resources:
    requests:
      storage: <Size>
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce`
