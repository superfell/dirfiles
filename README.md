# dirfiles

dirfiles is a tool to find directories with entries that match one regex, but not another, e.g. you can use this to find directories that contain a "cover.png" file, but no "cover.jpg" file.

You can almost certainly do this in pure bash, but for me this was easier.

create_jpeg.sh is a helper script to create a cover.jpg file from a supplied cover.png file

Use the 2 together to create cover.jpg files required by your media streamer for directories where you only have a .png version [you know, because it 2016, and why would any new product support png?, sigh], e.g.

./dirfiles -d /Volumes/Zaphod/Music/ | ./create_jpeg.sh
