/tmp/livego &
ffmpeg -stream_loop -1 -i /tmp/out.mp4 -f flv "rtmp://a.rtmp.youtube.com/live2/yp6r-myq6-z7t5-aht2"
# ffmpeg -stream_loop -1 -i /tmp/out.mp4 -f flv "rtmp://localhost:1935/live/pika"
# rtmp://a.rtmp.youtube.com/live2/yp6r-myq6-z7t5-aht2