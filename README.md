[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# DeepSort
🧠 AI powered image tagger backed by DeepDetect

# Prerequisites & installation

You need Docker installed, and you need to pull the DeepDetect image.
```
docker pull beniz/deepdetect_cpu
```

First download the latest release from https://github.com/CorentinB/DeepSort/releases
Make it executable with:
```
chmod +x deepsort
```

# Usage

Just input a folder, and all files will begin to be renames all the pictures files the following way:
```
identified-image-class_hash.ext
```
To start the tagging:
```
./deepsort FOLDER 
```

# (Really really quick) Benchmark

Tested on 605 files, it took 11m18s on an i7 7700K.
