#!/bin/bash

echo "--------------------"
echo "--------------------"
echo "   I AM RUNNING "
echo "--------------------"
echo "--------------------"

# Remove Current Assets Folder Including All Previous Modifications
rm -rf code/assets/
# Utilize Backup Copy.
cp -r code/assets_originals/ code/assets/

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