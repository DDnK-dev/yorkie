## Install Yorkie
# install istio with istio-operator
$ curl -L https://istio.io/downloadIstio | sh -
$ export PATH="$PATH:/home/dongwoo/istio-1.22.1/bin"
$ istioctl install -f <(curl -s https://raw.githubusercontent.com/yorkie-team/yorkie/main/build/charts/yorkie-cluster/istio-operator.yaml)
y

# 
# 