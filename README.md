# Check CM Cloud Downloaded Images
This program checks which image has been downloaded from CM Cloud (written in Go).

## Why did I make this program
CM Cloud is shutting down their service, so I have to download all my backed-up images into my local drive (I have like over 15000 images lol). Unfortunately, there was a network interruption during my downloading process. My program stopped downloading the images. When I resumed the download, it skipped thousands of images and started to download from a wrong image number. Having to check which image did it download or not from a list of over 15000 images is hideous. This is why I created this program for convenience.

## Setup
Assume that you have already followed the steps from this website: https://cloud.cmcm.com/p/offline/download_description.html, and have already started downloading the images. Now follow the steps as listed below:

1. Download this project's zip file to some other directory
2. Copy the following files in the directory that you are downloading your images to: 
    - cloud_gallery_urls.txt
    - cloud_gallery_urls.txt.done
    
    Paste them into this project's "files" folder (or replace them if these files already exist)
3. Run main.exe (can be done by double-clicking main.exe)

Note: if your downloaded images are being updated, remember to copy and paste the above 2 files again before running main.exe to ensure accurate results. Otherwise, you are good to just run main.exe after first-time setup.

## Example Output
Images downloaded from CM Cloud are all named as "someNumber.jpg". 
For example: `00000033.jpg`. This program prints a list of downloaded images in form of their numbers. They are demonstrated using ranges for readability.

The following shows an example of running main.exe:

```
Total number of images: 44

Number of Images Downloaded: 30

Downloaded images are listed as follows:
1 - 6
14
16 - 34
38 - 40
44
```
