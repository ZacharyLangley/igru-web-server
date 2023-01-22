#!/bin/bash
. /home/$USER/.nvm/nvm.sh
echo "Running Buf Generate..."
npm run buf:generate
echo "Cleaning up..."
sudo rm -rf ./apis
echo "Generation Complete. Don't lose your mind too quickly!"