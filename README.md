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

