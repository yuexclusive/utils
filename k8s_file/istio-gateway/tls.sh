kubectl create secret tls test-credential -n istio-system \
  --cert=/Users/yu/docker_volumes/tls/test.example.com.crt \
  --key=/Users/yu/docker_volumes/tls/test.example.com_key.key