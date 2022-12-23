#!/bin/sh

sshpass -p "$SSH_PASS" rsync -avz -e 'ssh -o StrictHostKeyChecking=no -p 22' --progress * abodsakka@abodsakka.xyz:/home/abodsakka/BTH-appen
