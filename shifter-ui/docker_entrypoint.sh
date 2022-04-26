#!/bin/bash

# Remove Current Assets Folder Including All Previous Modifications
rm -rf assets/
# Utilize Backup Copy.
cp -r assets_originals/ assets/

########################################
# Replace all References to Environment Variables  
# within JS Files.
########################################
export EXISTING_VARS=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,); 
for file in $JSFOLDER;
do
  cat $file | envsubst $EXISTING_VARS | tee $file
done



########################################
# Start NGINX Server
########################################
nginx