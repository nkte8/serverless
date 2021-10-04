# https://github.com/yt-dlp/yt-dlp/blob/master/yt_dlp/YoutubeDL.py
import yt_dlp as downloader
import os,json,base64,requests,sys,glob

FMTSET = { 
    "mp4" : ["mp4","m4a"],
    "webm": ["webm","webm"]
}
def dl_best_mp4(info_dict,fmt):
    # 品質リストのチェック
    # リストの下の方が品質が良いと仮定して取得
    best_video = {"format":"",'outtmpl': info_dict["title"] + '.video'}
    best_audio = {"format":"",'outtmpl': info_dict["title"] + '.audio'}
    for dlfmt in info_dict["formats"]:
        if (dlfmt.get("video_ext") == FMTSET.get(fmt)[0]):
            best_video["format"] = dlfmt.get("format_id")
        if (dlfmt.get("audio_ext") == FMTSET.get(fmt)[1]):
            best_audio["format"] = dlfmt.get("format_id")
    return best_video,best_audio

def stat_to_json(fp: str) -> dict:
    s_obj = os.stat(fp)
    return {k: getattr(s_obj, k) for k in dir(s_obj) if k.startswith('st_')}

if __name__ == "__main__":
    REST_URL=os.environ['REST_URL']
    response = requests.get(REST_URL.rstrip('/') + "/pop?type=dl")
    res = json.loads(response.text)

    ydl = downloader.YoutubeDL()
    with ydl:
        info_dict = ydl.extract_info(res.get("url"), download=False)
        _ = ydl.list_formats(info_dict)
        info_dict["title"] = \
            info_dict.get("title").replace('/', '／').replace('\\', '＼') + \
            ' [' + info_dict["id"] + ']'
        video,audio = dl_best_mp4(info_dict,res.get("fmt"))
        if video.get("format") != "": 
            ydl = downloader.YoutubeDL(video)
            _ = ydl.extract_info(res.get("url"), download=True)
        if audio.get("format") != "": 
            ydl = downloader.YoutubeDL(audio)
            _ = ydl.extract_info(res.get("url"), download=True)
        else:
            fname=info_dict.get("title") + "." + res.get("fmt")
            os.rename(info_dict.get("title") + ".video", fname)
            p_info=stat_to_json("../")
            os.chown(fname,uid=p_info.get("st_uid"),gid=p_info.get("st_gid"))
    ret = len(glob.glob(glob.escape(info_dict.get("title"))+"*")) > 0
    if (audio.get("format") != "") & ret:
        _ = requests.post(REST_URL.rstrip('/') + "/request",
            data={'url': base64.b64encode(info_dict.get("title").encode()),
                'type':'cnv','fmt':res.get("fmt"),'crf':res.get("crf")})
    sys.exit(not bool(ret)) # sysではTrue=0だが、pythonではTrue=1のため
