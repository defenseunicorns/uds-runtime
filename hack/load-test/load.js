import { K8s, kind } from 'kubernetes-fluent-client'

const namespace = 'load-test-' + Math.floor(Math.random() * 1000)

K8s(kind.Namespace).Apply({
  metadata: {
    name: namespace,
  },
})

for (let i = 0; i < 100; i++) {
  // Create the Pod in the namespace
  K8s(kind.Pod).Apply({
    metadata: {
      name: 'nginx-load-' + Math.floor(Math.random() * 1000),
      namespace,
    },
    spec: {
      containers: [
        {
          name: 'my-container',
          image: 'nginx',
        },
      ],
    },
  })
}
