#!/bin/bash
data_dir=/data/asset

crate_dir(){
    wd=$1
    sudo mkdir -p $wd
    sudo chown nobody:nogroup $wd
    sudo chmod 777 $wd -R
    echo "Created " $wd
}

crate_dir $data_dir

wget https://archive.org/download/ElephantsDream/ed_1024_512kb.mp4 -O $data_dir/ed.mp4 
wget https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm -O $data_dir/flower.mp4
wget https://archive.org/download/BigBuckBunny_124/Content/big_buck_bunny_720p_surround.mp4 -O $data_dir/buck.mp4

