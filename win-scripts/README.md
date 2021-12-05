# Running scripts

### Adding youtube-dl docker image to command prompt as **youtube-dl** command.

- First you need to add **win-scripts** folder to your windows the environment. Then run it:

```bash
yt
```

- It will run the doskey script saver and then you can start using **youtube-dl** command. Then you can download videos with running:

```bash
youtube-dl https://youtu.be/dQw4w9WgXcQ
```

- Also you can see other possible command options with running:

```bash
youtube-dl
```

Thanks [youtube-dl](https://github.com/ytdl-org/youtube-dl) for this awesome project.


### Using ffmpeg

```bash
ffmpeg -i "video.mp4" -f mp3 -ab 192000 -vn music.mp3
```