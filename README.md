[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com) 

[![Build Status](https://travis-ci.org/CorentinB/DeepSort.svg?branch=master)](https://travis-ci.org/CorentinB/DeepSort) [![Go Report Card](https://goreportcard.com/badge/github.com/CorentinB/DeepSort)](https://goreportcard.com/report/github.com/CorentinB/DeepSort) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/c62d2294e151492da4792fcb63b71d05)](https://www.codacy.com/project/CorentinB/DeepSort/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=CorentinB/DeepSort&amp;utm_campaign=Badge_Grade_Dashboard)

# DeepSort
🧠 AI powered image tagger backed by DeepDetect

# Why?

Because sometimes, you have folders full of badly named pictures, and you want to be able to understand what you have in your hard drive.

# Prerequisites & installation

You need DeepDetect installed, the easiest way is using docker:
```
docker pull beniz/deepdetect_cpu
docker run -d -p 8080:8080 beniz/deepdetect_cpu
```

Right now, the only supported installation of DeepDetect that works with DeepSort is the deepdetect_cpu container,
because it contain the good path for the pre-installed `resnet-50` and `googlenet` models.

Then, download the latest DeepSort release from https://github.com/CorentinB/DeepSort/releases

Unzip your release, rename it `DeepSort` and make it executable with:
```
chmod +x DeepSort
```

# Usage

DeepSort support few different parameters, you're obliged to fill two of them:
`--url` or `-u` that correspond to the URL of your DeepDetect server.
`--input` or `-i` that correspond to your local folder full of images.

For more informations, refeer to the helper:
```
./DeepSort --help

[-u|--url] is required
usage: deepsort [-h|--help] -u|--url "<value>" -i|--input "<value>"
                [-o|--output "<value>"] [-n|--network (resnet-50|googlenet)]
                [-R|--recursive] [-j|--jobs <integer>] [-d|--dry-run]

                AI powered image tagger backed by DeepDetect

Arguments:

  -h  --help       Print help information
  -u  --url        URL of your DeepDetect instance (i.e: http://localhost:8080)
  -i  --input      Your input folder.
  -o  --output     Your output folder, if output is set, original files will
                   not be renamed, but the renamed version will be copied in
                   the output folder.
  -n  --network    The pre-trained deep neural network you want to use, can be
                   resnet-50 or googlenet. Default: resnet-50
  -R  --recursive  Process files recursively.
  -j  --jobs       Number of parallel jobs. Default: 1
  -d  --dry-run    Just classify images and return results, do not apply.
```

# Todo list

- [X] Getting docker out of the loop (each user install his own DeepDetect)
- [X] ResNet 50 integration
- [X] Output folder (copy and not rename)
- [ ] NSFW tagging (Yahoo open_nsfw)
- [ ] XMP metadata writing
- [ ] GPU support
