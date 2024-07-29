import { K8s, kind } from 'kubernetes-fluent-client'

const namespace = 'load-test' + Math.floor(Math.random() * 1000);

async function createNamespace() {
  try {
    await K8s(kind.Namespace).Apply({
      metadata: {
        name: namespace,
      },
    });
    console.log(`Namespace ${namespace} created`);
  } catch (error) {
    console.error(`Error creating namespace ${namespace}:`, error);
  }
}

async function createPod(index) {
  try {
    const podName = 'nginx-load-' + Math.floor(Math.random() * 1000);
    await K8s(kind.Pod).Apply({
      metadata: {
        name: podName,
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
    });
    console.log(`Pod ${podName} created`);
  } catch (error) {
    console.error(`Error creating pod ${index}:`, error);
  }
}

async function createPods() {
  await createNamespace();
  for (let i = 0; i < 100; i++) {
    await createPod(i);
  }
}

createPods().then(() => {
  console.log('All pods created');
}).catch((error) => {
  console.error('Error creating pods:', error);
});
