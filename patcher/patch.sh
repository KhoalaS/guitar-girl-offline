#!/bin/bash
set -e
# unzip the split apk
unzip -d app app.apkm
# extract the global-metadata.dat from the assets apk into the current working directory
unzip -j app/split_base_assets.apk assets/bin/Data/Managed/Metadata/global-metadata.dat

cd il2cpp-stringliteral-patcher
python extract.py -i ../global-metadata.dat -o ../strings.json
cd ..

# extract the string value for the default game server URL and patch it replacing it with the new URL
cat strings.json | \
jq -c 'map(select(.value=="https://game.gtgl.pmang.cloud"))' | \
jq --arg API_URL "$GAME_API_URL" '.[0].value = $API_URL' > patched_strings.json

# patch the global-metadata.dat file
cd il2cpp-stringliteral-patcher
python patch.py -i ../global-metadata.dat -p ../patched_strings.json -o /tmp/global-metadata.dat
cd ..

cd app
# recreate folder structure of global-metadata.dat folder inside assets apk
mkdir -p assets/bin/Data/Managed/Metadata
# copy the patched global-metadata.dat into the dir and update the file inside the assets apk
cp /tmp/global-metadata.dat ./assets/bin/Data/Managed/Metadata/
zip split_base_assets.apk ./assets/bin/Data/Managed/Metadata/global-metadata.dat

# sign apks using the debug key
ls *.apk | while read line; do apksigner sign --ks ~/.android/debug.keystore --ks-pass pass:android $line; done
rm *.idsig
rm -rf assets
zip ../patched_app.apkm ./*


