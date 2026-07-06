# how to use
# sh convert_video.sh a.mp4 test_dir 1920x1080

# arguments
# S1 : src file path
# S2 : dst directory path
# S3 : resolution e.g 1920x1080

# make dst directory
if [ ! -e $2/hls ]; then
    mkdir -p $2/hls
fi
if [ ! -e $2/thumbnail ]; then
    mkdir -p $2/thumbnail
fi

# process video
# Http live streaming
# https://trac.ffmpeg.org/wiki/Encode/H.264
ffmpeg -i $1 -s $3 -start_number 0 -hls_segment_size 2000000 -f hls $2/hls/video.m3u8

# https:// ffmpeg.org/ffmpeg-formats.html#toc-image2-2
# https:// ffmpeg.org/ffmpeg.html#toc-Video-Options
ffmpeg -i $2/hls/video.m3u8 -s $3 -vframes 1 -f image2 $2/thumbnail/thumbnail.jpg