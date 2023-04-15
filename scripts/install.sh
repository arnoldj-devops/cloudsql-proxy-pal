#!/bin/bash
if which cloud_sql_proxy &>/dev/null
then
{    
    which cloudsql-proxy-pal | xargs sudo rm -rf 
}
fi
# curl -s https://raw.githubusercontent.com/arnoldj-devops/cloudsql-proxy-pal/master/scripts/install_prerequisites.sh | bash
LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/arnoldj-devops/cloudsql-proxy-pal/releases/latest)
LATEST_VERSION_TAG=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
LATEST_VERSION="${LATEST_VERSION_TAG:1}"
ARTIFACT_URL=https://github.com/arnoldj-devops/cloudsql-proxy-pal/releases/download/${LATEST_VERSION_TAG}/cloudsql-proxy-pal_${LATEST_VERSION}_Linux_x86_64.tar.gz
wget -qc  $ARTIFACT_URL -P /tmp && sudo tar -xvf /tmp/cloudsql-proxy-pal_${LATEST_VERSION}_Linux_x86_64.tar.gz -C /usr/local/bin/ cloudsql-proxy-pal >/dev/null 2>&1
cloudsql-proxy-pal connect --help
