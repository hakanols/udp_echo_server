gcloud auth configure-docker

gcloud projects list
# Get PROJECT_ID


# LOCAL_DOCKER_TAG = Arbitrary tag name for your docker image
cd [PROJECT_PATH]
docker build --rm -t [LOCAL_DOCKER_TAG] .
docker images # Verify

# DOCKER_IMAGE = Arbitrary name for your image

docker tag [LOCAL_DOCKER_TAG] gcr.io/[PROJECT_ID]/[DOCKER_IMAGE]
docker push gcr.io/[PROJECT_ID]/[DOCKER_IMAGE]

gcloud compute images list
# Chose VM_IMAGE. I chose a LTS cos-cloud ( Container-Optimized OS) image "cos-85-13310-1041-28".

# INSTANCE_NAMES = Arbitrary name for your VM
gcloud compute instances create-with-container [INSTANCE_NAMES] --container-image=gcr.io/[PROJECT_ID]/[DOCKER_IMAGE] --image=[VM_IMAGE] --tags=UDP_ON_1234

gcloud compute firewall-rules create "udp_open_on_1234" --allow=udp:1234 --source-ranges="0.0.0.0/0" --source-tags=UDP_ON_1234



# TODO Remove me
gcloud compute --project "udp-echo-296609" ssh --zone europe-north1-a test