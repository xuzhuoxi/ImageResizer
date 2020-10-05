# 取当前目录
CRTDIR=$(cd `dirname $0`; pwd) 

${CRTDIR}/IconGen -size=32,64,128,256x256,512 -in=./icon/source -out=./icon/target -format=png -ratio=75
