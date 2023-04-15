#!/bin/bash
if which cloud_sql_proxy &>/dev/null
then
{    
    echo "cloud_sql_proxy is installed"
}
else
{    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    {   
        if which gcloud &>/dev/null; then   
            echo "gcloud is installed"
        else
        {    
            echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
            sudo apt install apt-transport-https ca-certificates gnupg
            curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
            sudo apt update && sudo apt install -y google-cloud-sdk
            gcloud auth application-default login
        }
        fi
        wget -q https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
    }
    elif [[ "$OSTYPE" == "darwin"* ]]; then
    {
        if which gcloud &>/dev/null; then  
            echo "gcloud is installed"
        else
        {    
            brew tap homebrew/cask
            brew install google-cloud-sdk --cask
            echo 'source "$(brew --prefix)/Caskroom/google-cloud-sdk/latest/google-cloud-sdk/path.bash.inc"' >> ~/.bashrc
            gcloud auth application-default login
        }
        fi
        if [ "${arch_name}" = "x86_64" ]; then
        {
            curl -s -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.darwin.amd64
        }
        else
        {
            curl -s -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.darwin.arm64
        }
        fi
    }
    fi
    chmod +x cloud_sql_proxy
    sudo mv cloud_sql_proxy /usr/local/bin/
}
fi
