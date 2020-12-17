Run container in Google Cloud
=============================
This fiel contain instructions on how to push a and start a container to Goolge cloud.
There are att this time 3 ways to run a container in Google cloud:
* Kubernetes
* Run
* WM

In my case I just wanted to run a container. Uptime and so on was not as kritical.
Kubernetes is in my view the perferd way for any project.
But in my case where I just wanted to play aroun a bit that was noth worth the compexity that Kubernetes brings.
Google Run is a way to run containers as FaaS.
My echo server do not realy need a state but I found no way to get Run to accept a socket couse of its FaaS nature.
WM it is then. Found (Container-Optimized OS)[https://cloud.google.com/container-optimized-os/docs/concepts/features-and-benefits]
that was made for the task.

Instructions
============
Create a new project in Google console

Activate Container Registry API for that project

    gcloud projects list // Get PROJECT_ID
	gcloud config set project [PROJECT_ID]

DOCKER_IMAGE = Arbitrary name for your image

    cd [PROJECT_PATH]
    docker build --rm -t gcr.io/[PROJECT_ID]/[DOCKER_IMAGE] .
    docker images // Verify

    docker push gcr.io/[PROJECT_ID]/[DOCKER_IMAGE]
	gcloud container images list --repository=gcr.io/[PROJECT_ID] // Verify

INSTANCE_NAMES = Arbitrary name for your VM

TAG_NAME = Arbitrary name. E.g. "udp-1234"

RULE_NAME = Arbitrary name. E.g. "udp-1234"

    gcloud compute instances create-with-container [INSTANCE_NAMES] --container-image=gcr.io/[PROJECT_ID]/[DOCKER_IMAGE] --zone=europe-north1-a
	gcloud compute instances add-tags [INSTANCE_NAMES] --tags=[TAG_NAME]
    gcloud compute firewall-rules create [RULE_NAME] --allow=udp:1234 --source-ranges="0.0.0.0/0" --target-tags=[TAG_NAME]

    gcloud compute instances list // Get [EXTERNAL_IP]

