[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com) 

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c62d2294e151492da4792fcb63b71d05)](https://www.codacy.com/project/CorentinB/DeepSort/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=CorentinB/DeepSort&amp;utm_campaign=Badge_Grade_Dashboard) [![Go Report Card](https://goreportcard.com/badge/github.com/CorentinB/DeepSort)](https://goreportcard.com/report/github.com/CorentinB/DeepSort)

# DeepSort
🧠 AI powered image tagger backed by DeepDetect

# Why?

Because sometimes, you have folders full of badly named pictures, and you want to be able to understand what you have in your hard drive.

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

You also need your local 8080 port to not be mapped already.

# Usage

Just input a folder, and it will recursively rename all the pictures the following way:
```
identified-image-class_hash.ext
```
To start the tagging:
```
./deepsort FOLDER 
```

# (Really really quick) Benchmark

Tested on 605 files, it took 11m18s on an i7 7700K.

# Todo list

- [ ] Getting docker out of the loop (each user install his own DeepDetect)
- [ ] ResNet 50 integration
- [ ] XMP metadata writing
