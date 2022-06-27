CRTDIR=$(cd `dirname $0`; pwd)
# 单个产出
${CRTDIR}/ImageResizer_win_amd64.exe -mode=scale -src=src/SrcIcon.png -tar_file=tar/scale/TarIcon.png -scale=0.8 -format=jpg -ratio=85
# 批量产出
${CRTDIR}/ImageResizer_win_amd64.exe -mode=scale -include=png,jpg -src=src/dir,src/SrcIcon.png -tar_dir=tar/scale -scale=0.6,1,1.2 -ratio=85
