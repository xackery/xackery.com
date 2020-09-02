---
banner: "post/images/tech/cloud-storage.png"
categories:  ["tech"]
date:  "2017-07-10T00:52:46-07:00"
description:  "How to get started playing Karma, a modpack for Minecraft"
images:  ["post/images/karma/karma-day-10.png"]
menuTitle: "How I do backups"
tags:  ["tech"]
title:  "How I do backups"

---
I had a conversation with [PhotoJoseph (Youtuber)](https://www.youtube.com/photojoseph) during his live stream today, and felt inspired to write up a guide of how I do my backup solution. While it may be technical to some, the overall design is made to be low cost, versatile, easy once setup and scalable with at least some technical knowledge. There is the saying "You pay for convienence", but that isn't always the case.
<!--more-->

## The Scenario
Backups are important to a photographer. With the size of media in raw formats being so high, backup solutions can become a burden to maintain day to day. Disaster recovery is a pivotal part of any digital media user, and to accomplish this, it is typical to back up to multiple locations and multiple ways based on needs.

A generic backup plan may/should include these 5 steps:

1. <b>local</b> This is stored on a local computer and is typically the original copy.
2. <b>external backup</b> This is a physical external drive that is a mirror of the local copy.
3. <b>shared drive</b> This is a directory or drive shared via a network but is not always on the cloud. Dropbox is an example of this.
4. <b>cloud storage</b> This is made to be available remotely, and accessible from any location, and also acts as another backup.
5. <b>LTS cloud</b> LTS stands for long term storage, and is rarely accessed except in emergencies.

Over time, the local copy, external backup, and shared drive are likely to run out of space. What do you do when this happens? How do you maintain your backup and ensure data is available on all platforms as much as possible?


## My setup

* (external backup) [rsync](https://developer.apple.com/legacy/library/documentation/Darwin/Reference/ManPages/man1/rsync.1.html) is built into OSX.</a>
* (external backup) [Time Machine](https://support.apple.com/en-us/HT201250) is built into OSX.</b> Cost: Free
* <b>(shared drive) [Google Photos](https://photos.google.com/) for shared drive backup/distribution.</b></b> Cost: Free (for resized images)
* <b>(LTS cloud) (Coldline from Google Stoage)[https://cloud.google.com/storage/pricing] Coldline from Google Storage</a> for remote cloud backup storage.</b> Cost: $0.007/GB/mo.
* <b>(optional shared drive) (Syncthing)[https://syncthing.net/] for shared drive backup/distribution.</b> Cost: Free. Once they add a preview and not sync option to this, it will be viable for this setup.

### My breakdown of setup:

1. Take an older computer laying around and turn it into your server. Any OS works, really.
2. Purchase at least 4 external drives. If you need high performance you can opt in for internal drives, but be sure your computer supports hot-swappable SATA (or alternatively e-SATA). In my case, each drive is 4 terabytes.
3. Plug in each drive and label them, e.g. A, A-Backup, B, B-Backup. I give my drives unique labels to easily identify them. You can also date them or however you like to label.
4. Take any older computer you have laying around and install the external drives.
5. Install [Desktop Uploader](https://photos.google.com/apps) to sync each drive to your google photos. (For OSX/Linux, you can use [gDrive](https://github.com/prasmussen/gdrive) to do automated CLI interactions/rsync style backup). By default, RAW and other unsupported formats are ignored, but this is primarily for tracking and sharing with customers.
6. (technical, dropbox is an alternative) Configure sharing (samba or windows network sharing in any form) on the server of each letter drive, so users can access the drives.
7. (technical, dropbox is an alternative) Configure VPN to allow your sharing to be accessible from remote devices.
8. Create crons to rsync A to A-Backup (and all lettered drives), and email on failure.
9. Create another cron to rsync A to google cloud bucket backup, and email on failure.
10. Google photos can be set up to create albums based on folder. Once configured, you can share with customers specific albums by email for them to preview and give feedback.


Over time, when a drive gets too full of content, you can opt to purchase additional drives or empty them to re-use it as a new label. Group projects based on the different drives.

If a failure happens, you will have a shared backup, a external backup, on site, and a remote backup on the cloud via coldline.