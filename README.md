[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com) 

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c62d2294e151492da4792fcb63b71d05)](https://www.codacy.com/project/CorentinB/DeepSort/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=CorentinB/DeepSort&amp;utm_campaign=Badge_Grade_Dashboard) [![Go Report Card](https://goreportcard.com/badge/github.com/CorentinB/DeepSort)](https://goreportcard.com/report/github.com/CorentinB/DeepSort)

# DeepSort
🧠 AI powered image tagger backed by DeepDetect

# Why?

Because sometimes, you have folders full of badly named pictures, and you want to be able to understand what you have in your hard drive.

# Prerequisites & installation

You need DeepDetect installed, the easiest way is using docker:
```
docker pull beniz/deepdetect_cpu
docker run -d -p 8080:8080 -v /path/to/images:/path/to/images beniz/deepdetect_cpu
```

PLEASE NOTE THAT THE PATH IN THE HOST SHOULD BE THE SAME IN THE CONTAINER!
Example:
```
docker run -d -p 8080:8080 -v /home/corentin/Images:/home/corentin/Images beniz/deepdetect_cpu
```

If you prefeer using DeepDetect without Docker, refeer to the official documentation here:
https://github.com/jolibrain/deepdetect/blob/master/README.md
You'll find how to install it without Docker.

Then, download the latest DeepSort release from https://github.com/CorentinB/DeepSort/releases
Unzip your release, rename it `deepsort` and make it executable with:
```
chmod +x deepsort
```

# Usage

Right now, DeepSort doesnt support a lot of different parameters, you're obliged to fill two of them:
`--url` or `-u` that correspond to the URL of your DeepDetect server.
`--input` or `-i` that correspond to your local folder full of images.

For more informations, refeer to the helper:
```./deepsort --help
usage: deepsort [-h|--help] -u|--url "<value>" -i|--input "<value>"
                AI powered image tagger backed by DeepDetect
Arguments:

  -h  --help   Print help information
  -u  --url    URL of your DeepDetect instance (i.e: http://localhost:8080)
  -i  --input  Your input folder.
```

# (Really really quick) Benchmark

Tested on 605 files, it took 11m18s on an i7 7700K.

# Todo list

- [X] Getting docker out of the loop (each user install his own DeepDetect)
- [ ] Output folder (copy and not rename)
- [ ] ResNet 50 integration
- [ ] XMP metadata writing
- [ ] GPU support
