# 单个产出
ImageResizer_win_amd64.exe -mode=size -src=src/SrcIcon.png -tar_file=tar/size/TarIcon -size=512 -format=jpg -ratio=85
# 批量产出
ImageResizer_win_amd64.exe -mode=size -include=png,jpg -src=src/dir,src/SrcIcon.png -tar_dir=tar/size -size=512,256x256
