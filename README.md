PROJECT_ID = e.g "echo_server_1"

    gcloud projects create [PROJECT_ID]
    gcloud projects list // Verify

LOCAL_DOCKER_TAG = Arbitrary tag name for your docker image

    cd [PROJECT_PATH]
    docker build --rm -t [LOCAL_DOCKER_TAG] .
    docker images // Verify

DOCKER_IMAGE = Arbitrary name for your image

    docker tag [LOCAL_DOCKER_TAG] gcr.io/[PROJECT_ID]/[DOCKER_IMAGE]
    docker push gcr.io/[PROJECT_ID]/[DOCKER_IMAGE]
    gcloud compute images list // Verify
    
Chose VM_IMAGE. I chose a LTS cos-cloud ( Container-Optimized OS) image "cos-85-13310-1041-28".

INSTANCE_NAMES = Arbitrary name for your VM

    gcloud compute instances create-with-container [INSTANCE_NAMES] --container-image=gcr.io/[PROJECT_ID]/[DOCKER_IMAGE] --image=[VM_IMAGE] --tags=UDP_ON_1234
    gcloud compute firewall-rules create "udp_open_on_1234" --allow=udp:1234 --source-ranges="0.0.0.0/0" --source-tags=UDP_ON_1234

    gcloud compute instances list // Get [EXTERNAL_IP]
